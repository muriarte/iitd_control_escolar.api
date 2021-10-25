<script lang="ts">
  import IconButton, { Icon } from "@smui/icon-button/styled";
  import Svg from "@smui/common/Svg.svelte";
  import { mdiPlus, mdiTrashCanOutline, mdiPencil } from "@mdi/js";
  import { Label } from "@smui/common/styled";
  import Button from "@smui/button/styled";
  import Dialog, { Header, Title, Content, Actions } from "@smui/dialog/styled";
  import DataTable, { Head, Body, Row, Cell, SortValue } from "@smui/data-table/styled";
  import MateriaDetail from "./materiadetail.svelte";
  import type { Materia } from "../models/materia";
  import axios from "axios/dist/axios";

  let urlpathprefix: string = location.protocol + "//" + location.host + "/v1/";
  const httpClient = axios.create({
    timeout: 180000,
    baseURL: urlpathprefix,
  });

  let dialogOpen: boolean = false;

  let materia: Materia;
  let materias: Materia[] = [];
  let sort: keyof Materia = "id";
  // let sortDirection: Lowercase<keyof typeof SortValue> = "ascending";
  let sortDirection: string | number | symbol = "ascending";

  let displayError: string = "";

  cargaMaterias();

  function closeDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "reject":
        displayError = "Cancelado por el usuario.";
        break;
      case "accept":
        displayError = "Grabando!";
        grabaMateria(materia);
        displayError = "";
        break;
      default:
        // This means the user clicked the scrim or pressed Esc to close the dialog.
        // The actions will be "close".
        displayError = "";
        break;
    }
  }

  function handleSort() {
    materias.sort((a, b) => {
      const [aVal, bVal] = [a[sort], b[sort]][sortDirection === "ascending" ? "slice" : "reverse"]();
      if (typeof aVal === "string" && typeof bVal === "string") {
        return aVal.localeCompare(bVal);
      }
      return Number(aVal) - Number(bVal);
    });
    materias = materias;
  }

  function cargaMaterias() {
    displayError = "Por favor espere...";
    // Vaciamos el array
    materias.splice(0, materias.length);

    console.log("Cargando lista de materias... ");

    httpClient
      .get("materias")
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty materias list")
          return
        }
        let item: Materia;
        for (item of response.data.data) {
          item = corrigeFechas(item);
          materias.push(item);
        }
        materias = materias;

        console.log("Lista de materias:");
        console.log(materias);
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function corrigeFechas(item: Materia): Materia {
    // if (item.nacimiento && item.nacimiento.length > 10) {
    //   item.nacimiento = item.nacimiento.substr(0, 10);
    // }
    return item;
  }

  function nuevaMateria() {
    materia = {
      id: 0,
      nombre: "",
      observaciones: "",
      activo: "S",
    };
    dialogOpen = true;
  }

  function editMateria(id: number) {
    materia = materias.find((e) => e.id === id, "");
    console.log(materia);
    dialogOpen = true;
  }

  function grabaMateria(est: Materia) {
    console.log("Grabando materia:" + JSON.stringify(est));

    httpClient
      .post("materias", est)
      .then((response) => {
        displayError = "";
        console.log(response.data);
        est = response.data.data;
        est = corrigeFechas(est);

        // Actualiza el elemento en la lista de materias
        let found: boolean = false;
        for (let i: number = 0; i < materias.length; i++) {
          if (materias[i].id === est.id) {
            materias[i] = est;
            materias = materias;
            found = true;
            break;
          }
        }
        if (!found) {
          materias.push(est);
          materias = materias;
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function removeMateria(id: number) {
    console.log("Eliminando materia id:" + id);

    httpClient
      .delete("materias/" + id)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Si el status de retorno es falso no se eliminó la materia
        if (!response.data.data.status) return;

        // Elimina el elemento en la lista de materias
        for (let i: number = 0; i < materias.length; i++) {
          if (materias[i].id === id) {
            materias.splice(i, 1);
            materias = materias;
            break;
          }
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }
</script>

<h3>Catálogo de materias</h3>
<Button
  on:click={() => {
    nuevaMateria();
  }}
  variant="raised"
>
  <Label>Nueva materia</Label>
</Button>

<span class="dispErr">{displayError}</span>

<DataTable
  sortable
  bind:sort
  bind:sortDirection
  on:MDCDataTable:sorted={handleSort}
  table$aria-label="User list"
  style="width: 100%;"
>
  <Head>
    <Row>
      <!--
        Note: whatever you supply to "columnId" is
        appended with "-status-label" and used as an ID
        for the hidden label that describes the sort
        status to screen readers.
 
        You can localize those labels with the
        "sortAscendingAriaLabel" and
        "sortDescendingAriaLabel" props on the DataTable.
      -->
      <Cell />
      <Cell />
      <Cell numeric columnId="id">
        <!-- For numeric columns, icon comes first. -->
        <IconButton class="material-icons">arrow_upward</IconButton>
        <Label>ID</Label>
      </Cell>
      <Cell columnId="nombre" style="width: 100%;">
        <Label>Nombre</Label>
        <!-- For non-numeric columns, icon comes second. -->
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="observaciones">
        <Label>Observaciones</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="activo">
        <Label>Activo</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
    </Row>
  </Head>
  <Body>
    {#each materias as item (item.id)}
      <Row>
        <Cell>
          <a href="#edit" on:click|preventDefault={() => editMateria(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="blue" fill="currentColor" d={mdiPencil} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#delete" on:click|preventDefault={() => removeMateria(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiTrashCanOutline} />
            </Icon></a
          >
        </Cell>
        <Cell numeric>{item.id}</Cell>
        <Cell>{item.nombre}</Cell>
        <Cell>{item.observaciones}</Cell>
        <Cell>{item.activo}</Cell>
      </Row>
    {/each}
  </Body>
</DataTable>

<Dialog
  bind:open={dialogOpen}
  fullscreen
  aria-labelledby="fullscreen-title"
  aria-describedby="fullscreen-content"
  on:MDCDialog:closed={closeDialogHandler}
>
  <Header>
    <Title id="fullscreen-title">Detalles de la materia</Title>
    <IconButton action="close" class="material-icons">close</IconButton>
  </Header>
  <Content id="fullscreen-content">
    <MateriaDetail bind:materia />
  </Content>
  <Actions>
    <Button action="reject">
      <Label>Reject</Label>
    </Button>
    <Button action="accept" defaultAction>
      <Label>Accept</Label>
    </Button>
  </Actions>
</Dialog>

<style>
  .dispErr {
    color: red;
  }
</style>
