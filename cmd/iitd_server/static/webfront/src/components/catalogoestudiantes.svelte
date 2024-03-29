<script lang="ts">
  import IconButton, { Icon } from "@smui/icon-button/styled";
  import Svg from "@smui/common/elements/Svg.svelte";
  import { mdiPlus, mdiTrashCanOutline, mdiPencil, mdiBook, mdiNoteTextOutline } from "@mdi/js";
  import { Label } from "@smui/common/styled";
  import Button from "@smui/button/styled";
  import Dialog, { Header, Title, Content, Actions } from "@smui/dialog/styled";
  import DataTable, { Head, Body, Row, Cell, SortValue } from "@smui/data-table/styled";
  import EstudianteDetail from "./estudiantedetail.svelte";
  import type { Estudiante } from "../models/estudiante";
  import axios from "axios/dist/axios";

  import { useLocation, useNavigate } from "svelte-navigator";
  const navigate = useNavigate();
  const loc = useLocation();
  $: console.log("catalogoestudiantes");
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

  let estudiante: Estudiante;
  let estudiantes: Estudiante[] = [];
  let sort: keyof Estudiante = "id";
  // let sortDirection: Lowercase<keyof typeof SortValue> = "ascending";
  let sortDirection: "ascending" | "descending" | "none" | "other" = "ascending";

  let displayError: string = "";

  cargaEstudiantes();

  function closeDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "cancelar":
        displayError = "Cancelado por el usuario.";
        break;
      case "aceptar":
        displayError = "Grabando!";
        if (grabaEstudiante(estudiante)) {
          displayError = "";
        }
        break;
      default:
        // This means the user clicked the scrim or pressed Esc to close the dialog.
        // The actions will be "close".
        displayError = "";
        break;
    }
  }

  function handleSort() {
    estudiantes.sort((a, b) => {
      const [aVal, bVal] = [a[sort], b[sort]][sortDirection === "ascending" ? "slice" : "reverse"]();
      if (typeof aVal === "string" && typeof bVal === "string") {
        return aVal.localeCompare(bVal);
      }
      return Number(aVal) - Number(bVal);
    });
    estudiantes = estudiantes;
  }

  function cargaEstudiantes(): boolean {
    displayError = "Por favor espere...";
    // Vaciamos el array
    estudiantes.splice(0, estudiantes.length);

    console.log("Cargando lista de estudiantes... ");

    httpClient
      .get("students")
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty students list");
          return;
        }
        let item: Estudiante;
        for (item of response.data.data) {
          item = corrigeFechas(item);
          estudiantes.push(item);
        }
        estudiantes = estudiantes;

        console.log("Lista de estudiantes:");
        console.log(estudiantes);
      })
      .catch((error) => {
        displayError = error;
        return false;
      });
    return true;
  }

  function corrigeFechas(item: Estudiante): Estudiante {
    if (item.nacimiento && item.nacimiento.length > 10) {
      item.nacimiento = item.nacimiento.substr(0, 10);
    }
    if (item.fechaInicio && item.fechaInicio.length > 10) {
      item.fechaInicio = item.fechaInicio.substr(0, 10);
    }
    return item;
  }

  function nuevoEstudiante() {
    estudiante = {
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

  function editEstudiante(id: number) {
    estudiante = estudiantes.find((e) => e.id === id, "");
    console.log(estudiante);
    dialogOpen = true;
  }

  function showMaterias(estudianteId: number) {
    navigate("/estudiantematerias/" + estudianteId, {
      replace: false,
      state: {},
    });
  }

  function showObservaciones(estudianteId: number) {
    navigate("/estudianteobs/" + estudianteId, {
      replace: false,
      state: {},
    });
  }

  function grabaEstudiante(est: Estudiante): boolean {
    console.log("Grabando estudiante:" + JSON.stringify(est));

    httpClient
      .post("students", est)
      .then((response) => {
        displayError = "";
        console.log(response.data);
        est = response.data.data;
        est = corrigeFechas(est);

        // Actualiza el elemento en la lista de estudiantes
        let found: boolean = false;
        for (let i: number = 0; i < estudiantes.length; i++) {
          if (estudiantes[i].id === est.id) {
            estudiantes[i] = est;
            estudiantes = estudiantes;
            found = true;
            break;
          }
        }
        if (!found) {
          estudiantes.push(est);
          estudiantes = estudiantes;
        }
      })
      .catch((error) => {
        displayError = error;
        return false;
      });
    return true;
  }

  // ELIMINAR Estudiante
  function eliminaEstudianteAsk(id: number, name: string) {
    idToRemove = id;
    yesnoMsg = id + "-" + name;
    yesnoDialogOpen = true;
  }

  function closeYesnoDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "Si":
        displayError = "Borrando!";
        eliminaEstudiante(idToRemove);
        break;
      case "No":
        displayError = "Cancelado por el usuario.";
        break;
    }
  }

  function eliminaEstudiante(id: number): boolean {
    console.log("Eliminando estudiante id:" + id);

    httpClient
      .delete("students/" + id)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Si el status de retorno es falso no se eliminó el estudiante
        if (!response.data.data.status) return;

        // Elimina el elemento en la lista de estudiantes
        for (let i: number = 0; i < estudiantes.length; i++) {
          if (estudiantes[i].id === id) {
            estudiantes.splice(i, 1);
            estudiantes = estudiantes;
            break;
          }
        }
      })
      .catch((error) => {
        displayError = error;
        return false;
      });
    return true;
  }
</script>

<h3>Catálogo de estudiantes</h3>
<Button
  on:click={() => {
    nuevoEstudiante();
  }}
  variant="raised"
>
  <Label>Nuevo estudiante</Label>
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
    {#each estudiantes as item (item.id)}
      <Row>
        <Cell>
          <a href="#edit" on:click|preventDefault={() => editEstudiante(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="blue" fill="currentColor" d={mdiPencil} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#delete" on:click|preventDefault={() => eliminaEstudianteAsk(item.id, item.apellidos + " / " + item.nombres)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiTrashCanOutline} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#materias" on:click|preventDefault={() => showMaterias(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiBook} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a href="#obs" on:click|preventDefault={() => showObservaciones(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiNoteTextOutline} />
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
    <Title id="fullscreen-title">Detalles del estudiante</Title>
    <IconButton action="close" class="material-icons">close</IconButton>
  </Header>
  <Content id="fullscreen-content">
    <EstudianteDetail bind:estudiante />
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
    <Title id="fullscreen-title">Eliminar estudiante</Title>
  </Header>
  <Content id="fullscreen-content">
    <Label>¿Desea eliminar el estudiante [{yesnoMsg}]?</Label>
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
