<script lang="ts">
  import IconButton, { Icon } from "@smui/icon-button/styled";
  import Svg from "@smui/common/Svg.svelte";
  import { mdiPlus, mdiTrashCanOutline, mdiPencil } from "@mdi/js";
  import { Label } from "@smui/common/styled";
  import Button from "@smui/button/styled";
  import Dialog, { Header, Title, Content, Actions } from "@smui/dialog/styled";
  import DataTable, { Head, Body, Row, Cell, SortValue } from "@smui/data-table/styled";
  import MaestroDetail from "./maestrodetail.svelte";
  import type { Maestro } from "../models/maestro";
  import axios from "axios/dist/axios";

  let urlpathprefix: string = location.protocol + "//" + location.host + "/v1/";
  const httpClient = axios.create({
    timeout: 180000,
    baseURL: urlpathprefix,
  });

  let dialogOpen: boolean = false;

  let maestro: Maestro;
  let maestros: Maestro[] = [];
  let sort: keyof Maestro = "id";
  // let sortDirection: Lowercase<keyof typeof SortValue> = "ascending";
  let sortDirection: string | number | symbol = "ascending";

  let displayError: string = "";

  cargaMaestros();

  function closeDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "reject":
        displayError = "Cancelado por el usuario.";
        break;
      case "accept":
        displayError = "Grabando!";
        grabaMaestro(maestro);
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
    maestros.sort((a, b) => {
      const [aVal, bVal] = [a[sort], b[sort]][sortDirection === "ascending" ? "slice" : "reverse"]();
      if (typeof aVal === "string" && typeof bVal === "string") {
        return aVal.localeCompare(bVal);
      }
      return Number(aVal) - Number(bVal);
    });
    maestros = maestros;
  }

  function cargaMaestros() {
    displayError = "Por favor espere...";
    // Vaciamos el array
    maestros.splice(0, maestros.length);

    console.log("Cargando lista de maestros... ");

    httpClient
      .get("maestros")
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty maestros list")
          return
        }
        let item: Maestro;
        for (item of response.data.data) {
          item = corrigeFechas(item);
          maestros.push(item);
        }
        maestros = maestros;

        console.log("Lista de maestros:");
        console.log(maestros);
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function corrigeFechas(item: Maestro): Maestro {
    if (item.nacimiento && item.nacimiento.length > 10) {
      item.nacimiento = item.nacimiento.substr(0, 10);
    }
    if (item.fechaInicio && item.fechaInicio.length > 10) {
      item.fechaInicio = item.fechaInicio.substr(0, 10);
    }
    return item;
  }

  function nuevoMaestro() {
    maestro = {
      id: 0,
      nombres: "",
      apellidos: "",
      nacimiento: "2000-01-01",
      sexo: "",
      calle: "",
      numeroExt: "",
      numeroInt: "",
      colonia: "",
      municipio: "",
      estado: "",
      pais: "",
      cp: "",
      telCasa: "",
      telCelular: "",
      email: "",
      fechaInicio: "2000-01-01",
      observaciones: "",
      activo: "S",
    };
    dialogOpen = true;
  }

  function editMaestro(id: number) {
    maestro = maestros.find((e) => e.id === id, "");
    console.log(maestro);
    dialogOpen = true;
  }

  function grabaMaestro(est: Maestro) {
    console.log("Grabando maestro:" + JSON.stringify(est));

    httpClient
      .post("maestros", est)
      .then((response) => {
        displayError = "";
        console.log(response.data);
        est = response.data.data;
        est = corrigeFechas(est);

        // Actualiza el elemento en la lista de maestros
        let found: boolean = false;
        for (let i: number = 0; i < maestros.length; i++) {
          if (maestros[i].id === est.id) {
            maestros[i] = est;
            maestros = maestros;
            found = true;
            break;
          }
        }
        if (!found) {
          maestros.push(est);
          maestros = maestros;
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function removeMaestro(id: number) {
    console.log("Eliminando maestro id:" + id);

    httpClient
      .delete("maestros/" + id)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Si el status de retorno es falso no se eliminó el maestro
        if (!response.data.data.status) return;

        // Elimina el elemento en la lista de maestros
        for (let i: number = 0; i < maestros.length; i++) {
          if (maestros[i].id === id) {
            maestros.splice(i, 1);
            maestros = maestros;
            break;
          }
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }
</script>

<h3>Catálogo de maestros</h3>
<Button
  on:click={() => {
    nuevoMaestro();
  }}
  variant="raised"
>
  <Label>Nuevo maestro</Label>
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
      <Cell columnId="nombres" style="width: 100%;">
        <Label>Nombre(s)</Label>
        <!-- For non-numeric columns, icon comes second. -->
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="apellidos">
        <Label>Apellido(s)</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="email">
        <Label>Email</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <!-- You can turn off sorting for a column. -->
      <Cell columnId="fechaInicio">
        <IconButton class="material-icons">arrow_upward</IconButton>
        <Label>Fecha de Inicio</Label>
      </Cell>
    </Row>
  </Head>
  <Body>
    {#each maestros as item (item.id)}
      <Row>
        <Cell>
          <a href="#edit" on:click|preventDefault={() => editMaestro(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="blue" fill="currentColor" d={mdiPencil} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#delete" on:click|preventDefault={() => removeMaestro(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiTrashCanOutline} />
            </Icon></a
          >
        </Cell>
        <Cell numeric>{item.id}</Cell>
        <Cell>{item.nombres}</Cell>
        <Cell>{item.apellidos}</Cell>
        <Cell>{item.email}</Cell>
        <Cell>{item.fechaInicio}</Cell>
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
    <Title id="fullscreen-title">Detalles del maestro</Title>
    <IconButton action="close" class="material-icons">close</IconButton>
  </Header>
  <Content id="fullscreen-content">
    <MaestroDetail bind:maestro />
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
