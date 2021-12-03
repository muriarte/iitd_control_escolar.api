<script lang="ts">
  import IconButton, { Icon } from "@smui/icon-button/styled";
  import Svg from "@smui/common/elements/Svg.svelte";
  import { mdiPlus, mdiTrashCanOutline, mdiPencil, mdiBook } from "@mdi/js";
  import { Label } from "@smui/common/styled";
  import Button from "@smui/button/styled";
  import Dialog, { Header, Title, Content, Actions } from "@smui/dialog/styled";
  import DataTable, { Head, Body, Row, Cell, SortValue } from "@smui/data-table/styled";
  import Tooltip, { Wrapper } from "@smui/tooltip";
  import EstudianteObsDetail from "./estudianteobsdetail.svelte";
  import type { Estudiante } from "../models/estudiante";
  import { EstudianteObs } from "../models/estudianteobs";
  import axios from "axios/dist/axios";

  import { useLocation, useNavigate } from "svelte-navigator";
  import { debug } from "svelte/internal";

  export let studentId: number = 0;

  const navigate = useNavigate();
  const loc = useLocation();
  $: console.log("estudianteobss");
  $: console.log($loc);

  let urlpathprefix: string = location.protocol + "//" + location.host + "/v1/";
  const httpClient = axios.create({
    timeout: 180000,
    baseURL: urlpathprefix,
  });

  let dialogOpen: boolean = false;

  let yesnoDialogOpen: boolean = false;
  let yesnoMsg: string = "";
  let idToRemove: number = 0;

  let estudianteObs: EstudianteObs;
  let observaciones: EstudianteObs[] = [];
  let sort: keyof Estudiante = "id";
  // let sortDirection: Lowercase<keyof typeof SortValue> = "ascending";
  let sortDirection: "ascending" | "descending" | "none" | "other" = "ascending";

  let displayError: string = "";
  let clonedEstudianteObs: EstudianteObs = EstudianteObs.newEmpty();

  cargaEstudianteObs(studentId);

  function closeDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "cancelar":
        displayError = "Cancelado por el usuario.";
        break;
      case "aceptar":
        displayError = "Grabando!";
        // Detectamos y corregimos un comportamiento erroneo que se
        // presenta cuando editamos por primera vez una materia del
        // alumno inmediatamente despues de desplegarse la lista de
        // materias del alumno (no pudimos encontrar la causa)
        // if (clonedEstudianteObs.materiaId === undefined) {
        //   var found: boolean = false;
        //   for (let i: number = 0; i < materias.length; i++) {
        //     if (materias[i].id === clonedEstudianteObs.id) {
        //       clonedEstudianteObs.materiaId = materias[i].materiaId;
        //       console.log(
        //         "Pudimos corregir caso de cuando el campo materiaId se vuelve undefined:" +
        //           clonedEstudianteObs.materiaId
        //       );
        //       found = true;
        //       break;
        //     }
        //   }
        //   if (!found) {
        //     console.log("No pudimos corregir caso de cuando el campo materiaId se vuelve undefined");
        //   }
        // }
        //Fin de deteccion y correccion de error

        grabaEstudianteObs(clonedEstudianteObs);
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
    observaciones.sort((a, b) => {
      const [aVal, bVal] = [a[sort], b[sort]][sortDirection === "ascending" ? "slice" : "reverse"]();
      if (typeof aVal === "string" && typeof bVal === "string") {
        return aVal.localeCompare(bVal);
      }
      return Number(aVal) - Number(bVal);
    });
    observaciones = observaciones;
  }

  function cargaEstudianteObs(studentId: number) {
    displayError = "Por favor espere...";
    // Vaciamos el array
    observaciones.splice(0, observaciones.length);

    console.log("Cargando lista de estudiantes... ");

    httpClient
      .get("studentobs?studentId=" + studentId)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty student materias list");
          return;
        }
        let item: EstudianteObs;
        for (item of response.data.data) {
          item = corrigeFechas(item);
          observaciones.push(item);
        }
        observaciones = observaciones;

        console.log("Lista de estudiante-materias:");
        console.log(observaciones);
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function corrigeFechas(item: EstudianteObs): EstudianteObs {
    if (item.fecha && item.fecha.length > 10) {
      item.fecha = item.fecha.substr(0, 10);
    }
    return item;
  }

  function nuevoEstudianteObs(studentId: number) {
    estudianteObs = new EstudianteObs(0, studentId, "2000-01-01", "");
    clonedEstudianteObs = EstudianteObs.guaranteedClone(estudianteObs);

    dialogOpen = true;
  }

  function editEstudianteObs(id: number) {
    estudianteObs = observaciones.find((e) => e.id === id, "");
    clonedEstudianteObs = EstudianteObs.guaranteedClone(estudianteObs);

    console.log(clonedEstudianteObs);
    dialogOpen = true;
  }

  function grabaEstudianteObs(est: EstudianteObs) {
    console.log("Grabando estudiante-materia:" + JSON.stringify(est));

    httpClient
      .post("studentobs", est)
      .then((response) => {
        displayError = "";
        console.log(response.data);
        est = response.data.data;
        est = corrigeFechas(est);

        // Actualiza el elemento en la lista de estudiantes
        let found: boolean = false;
        for (let i: number = 0; i < observaciones.length; i++) {
          if (observaciones[i].id === est.id) {
            observaciones[i] = est;
            observaciones = observaciones;
            found = true;
            break;
          }
        }
        if (!found) {
          observaciones.push(est);
          observaciones = observaciones;
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }

  // ELIMINAR EstudianteObs
  function eliminaEstudianteObsAsk(id: number) {
    idToRemove = id;
    yesnoMsg = id + "";
    yesnoDialogOpen = true;
  }

  function closeYesnoDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "Si":
        displayError = "Borrando!";
        eliminaEstudianteObs(idToRemove);
        break;
      case "No":
        displayError = "Cancelado por el usuario.";
        break;
    }
  }

  function eliminaEstudianteObs(id: number) {
    console.log("Eliminando observación de estudiante id:" + id);

    httpClient
      .delete("studentobs/" + id)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Si el status de retorno es falso no se eliminó el estudiante
        if (!response.data.data.status) return;

        // Elimina el elemento en la lista de estudiantes
        for (let i: number = 0; i < observaciones.length; i++) {
          if (observaciones[i].id === id) {
            observaciones.splice(i, 1);
            observaciones = observaciones;
            break;
          }
        }
      })
      .catch((error) => {
        displayError = error;
      });
  }
</script>

<!-- {@debug estudianteObs}; -->

<h3>Observaciones del estudiante</h3>
<Button
  on:click={() => {
    nuevoEstudianteObs(studentId);
  }}
  variant="raised"
>
  <Label>Agregar observación</Label>
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
      <!-- <Cell numeric columnId="id">
        < !-- For numeric columns, icon comes first. -- >
        <IconButton class="material-icons">arrow_upward</IconButton>
        <Label>ID</Label>
      </Cell>
      <Cell numeric columnId="materiaId">
        < !-- For numeric columns, icon comes first. -- >
        <IconButton class="material-icons">arrow_upward</IconButton>
        <Label>ID Materia</Label>
      </Cell> -->
      <Cell columnId="fecha">
        <Label>Fecha</Label>
        <!-- For non-numeric columns, icon comes second. -->
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="observacion" style="width: 100%;">
        <Label>Observación:</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
    </Row>
  </Head>
  <Body>
    {#each observaciones as item (item.id)}
      <Row>
        <Cell>
          <a href="#edit" on:click|preventDefault={() => editEstudianteObs(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="blue" fill="currentColor" d={mdiPencil} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#delete" on:click|preventDefault={() => eliminaEstudianteObsAsk(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiTrashCanOutline} />
            </Icon></a
          >
        </Cell>
        <Cell>{item.fecha}</Cell>
        <Cell>{item.observacion}</Cell>
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
    <Title id="fullscreen-title">Detalles del estudiante</Title>
    <IconButton action="close" class="material-icons">close</IconButton>
  </Header>
  <Content id="fullscreen-content">
    <EstudianteObsDetail bind:estudianteObs={clonedEstudianteObs} />
  </Content>
  <Actions>
    <Button action="cancelar">
      <Label>Cancelar</Label>
    </Button>
    <Button action="aceptar" defaultAction>
      <Label>Aceptar</Label>
    </Button>
  </Actions>
</Dialog>

<Dialog
  bind:open={yesnoDialogOpen}
  aria-labelledby="fullscreen-title"
  aria-describedby="fullscreen-content"
  on:MDCDialog:closed={closeYesnoDialogHandler}
>
  <Header>
    <Title id="fullscreen-title">Eliminar observación del estudiante</Title>
  </Header>
  <Content id="fullscreen-content">
    <Label>¿Desea eliminar la observación [{yesnoMsg}] del estudiante seleccionado?</Label>
  </Content>
  <Actions>
    <Button action="No" defaultAction>
      <Label>No</Label>
    </Button>
    <Button action="Si">
      <Label>Si</Label>
    </Button>
  </Actions>
</Dialog>

<style>
  .dispErr {
    color: red;
  }
</style>
