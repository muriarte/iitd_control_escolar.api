<script lang="ts">
  import IconButton, { Icon } from "@smui/icon-button/styled";
  import Svg from "@smui/common/elements/Svg.svelte";
  import { mdiPlus, mdiTrashCanOutline, mdiPencil, mdiBook } from "@mdi/js";
  import { Label } from "@smui/common/styled";
  import Button from "@smui/button/styled";
  import Dialog, { Header, Title, Content, Actions } from "@smui/dialog/styled";
  import DataTable, { Head, Body, Row, Cell, SortValue } from "@smui/data-table/styled";
  import EstudianteMateriaDetail from "./estudiantemateriadetail.svelte";
  import type { Estudiante } from "../models/estudiante";
  import { EstudianteMateria } from "../models/estudiantemateria";
  import axios from "axios/dist/axios";

  import { useLocation, useNavigate } from "svelte-navigator";
  import { debug } from "svelte/internal";

  export let studentId: number = 0;

  const navigate = useNavigate();
  const loc = useLocation();
  $: console.log("estudiantematerias");
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

  let estudianteMateria: EstudianteMateria;
  let materias: EstudianteMateria[] = [];
  let sort: keyof Estudiante = "id";
  // let sortDirection: Lowercase<keyof typeof SortValue> = "ascending";
  let sortDirection: string | number | symbol = "ascending";

  let displayError: string = "";
  let clonedEstudianteMateria: EstudianteMateria = EstudianteMateria.newEmpty();

  cargaEstudianteMaterias(studentId);

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
        if (clonedEstudianteMateria.materiaId === undefined) {
          var found: boolean = false;
          for (let i: number = 0; i < materias.length; i++) {
            if (materias[i].id === clonedEstudianteMateria.id) {
              clonedEstudianteMateria.materiaId = materias[i].materiaId;
              console.log(
                "Pudimos corregir caso de cuando el campo materiaId se vuelve undefined:" +
                  clonedEstudianteMateria.materiaId
              );
              found = true;
              break;
            }
          }
          if (!found) {
            console.log("No pudimos corregir caso de cuando el campo materiaId se vuelve undefined");
          }
        }
        //Fin de deteccion y correccion de error

        grabaEstudianteMateria(clonedEstudianteMateria);
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

  function cargaEstudianteMaterias(studentId: number) {
    displayError = "Por favor espere...";
    // Vaciamos el array
    materias.splice(0, materias.length);

    console.log("Cargando lista de estudiantes... ");

    httpClient
      .get("studentmaterias?studentId=" + studentId)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty student materias list");
          return;
        }
        let item: EstudianteMateria;
        for (item of response.data.data) {
          item = corrigeFechas(item);
          materias.push(item);
        }
        materias = materias;

        console.log("Lista de estudiante-materias:");
        console.log(materias);
      })
      .catch((error) => {
        displayError = error;
      });
  }

  function corrigeFechas(item: EstudianteMateria): EstudianteMateria {
    if (item.inicio && item.inicio.length > 10) {
      item.inicio = item.inicio.substr(0, 10);
    }
    if (item.fin && item.fin.length > 10) {
      item.fin = item.fin.substr(0, 10);
    }
    return item;
  }

  function nuevoEstudianteMateria(studentId: number) {
    estudianteMateria = new EstudianteMateria(0, studentId, 0, "", "2000-01-01", "2000-01-01", "");
    clonedEstudianteMateria = EstudianteMateria.guaranteedClone(estudianteMateria);

    dialogOpen = true;
  }

  function editEstudianteMateria(id: number) {
    estudianteMateria = materias.find((e) => e.id === id, "");
    clonedEstudianteMateria = EstudianteMateria.guaranteedClone(estudianteMateria);

    console.log(clonedEstudianteMateria);
    dialogOpen = true;
  }

  function grabaEstudianteMateria(est: EstudianteMateria) {
    console.log("Grabando estudiante-materia:" + JSON.stringify(est));

    httpClient
      .post("studentmaterias", est)
      .then((response) => {
        displayError = "";
        console.log(response.data);
        est = response.data.data;
        est = corrigeFechas(est);

        // Actualiza el elemento en la lista de estudiantes
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

  // ELIMINAR EstudianteMateria
  function eliminaEstudianteMateriaAsk(id: number, name: string) {
    idToRemove = id;
    yesnoMsg = id + "-" + name;
    yesnoDialogOpen = true;
  }

  function closeYesnoDialogHandler(e: CustomEvent<{ action: string }>) {
    switch (e.detail.action) {
      case "Si":
        displayError = "Borrando!";
        eliminaEstudianteMateria(idToRemove);
        break;
      case "No":
        displayError = "Cancelado por el usuario.";
        break;
    }
  }

  function eliminaEstudianteMateria(id: number) {
    console.log("Eliminando estudiante id:" + id);

    httpClient
      .delete("studentmaterias/" + id)
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Si el status de retorno es falso no se eliminó el estudiante
        if (!response.data.data.status) return;

        // Elimina el elemento en la lista de estudiantes
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

<!-- {@debug estudianteMateria}; -->

<h3>Materias del estudiante</h3>
<Button
  on:click={() => {
    nuevoEstudianteMateria(studentId);
  }}
  variant="raised"
>
  <Label>Agregar materia</Label>
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
      <Cell columnId="nombreMateria" style="width: 100%;">
        <Label>Nombre de la materia</Label>
        <!-- For non-numeric columns, icon comes second. -->
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="inicio">
        <Label>Fecha Inicio</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <Cell columnId="fin">
        <Label>Fecha fin</Label>
        <IconButton class="material-icons">arrow_upward</IconButton>
      </Cell>
      <!-- You can turn off sorting for a column. -->
      <Cell columnId="observaciones">
        <IconButton class="material-icons">arrow_upward</IconButton>
        <Label>Observaciones</Label>
      </Cell>
    </Row>
  </Head>
  <Body>
    {#each materias as item (item.id)}
      <Row>
        <Cell>
          <a href="#edit" on:click|preventDefault={() => editEstudianteMateria(item.id)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="blue" fill="currentColor" d={mdiPencil} />
            </Icon></a
          >
        </Cell>
        <Cell>
          <a
            href="#delete"
            on:click|preventDefault={() => eliminaEstudianteMateriaAsk(item.id, item.materiaNombre)}
            ><Icon component={Svg} viewBox="0 0 24 24" style="width:24px;height:24px;">
              <path Color="red" fill="currentColor" d={mdiTrashCanOutline} />
            </Icon></a
          >
        </Cell>
        <!-- <Cell numeric>{item.id}</Cell>
        <Cell>{item.materiaId}</Cell> -->
        <Cell>{item.materiaNombre}({item.id})({item.materiaId})</Cell>
        <Cell>{item.inicio}</Cell>
        <Cell>{item.fin}</Cell>
        <Cell>{item.observaciones}</Cell>
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
    <EstudianteMateriaDetail bind:estudianteMateria={clonedEstudianteMateria} />
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
    <Title id="fullscreen-title">Eliminar materia del estudiante</Title>
  </Header>
  <Content id="fullscreen-content">
    <Label>¿Desea eliminar la materia [{yesnoMsg}] del estudiante seleccionado?</Label>
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
