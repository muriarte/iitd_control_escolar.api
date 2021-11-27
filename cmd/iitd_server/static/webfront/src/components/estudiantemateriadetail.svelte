<script lang="ts">
  import Textfield from "@smui/textfield/styled";
  import LayoutGrid, { Cell } from "@smui/layout-grid/styled";
  import Radio from "@smui/radio/styled";
  import FormField from "@smui/form-field/styled";
  import { Label } from "@smui/common/styled";
  import Checkbox from "@smui/checkbox/styled";
  import Select, { Option } from "@smui/select/styled";
  import { EstudianteMateria } from "../models/estudiantemateria";
  import type { Materia } from "../models/materia";
  import axios from "axios/dist/axios";

  let urlpathprefix: string = location.protocol + "//" + location.host + "/v1/";
  const httpClient = axios.create({
    timeout: 180000,
    baseURL: urlpathprefix,
  });

  export let estudianteMateria: EstudianteMateria = EstudianteMateria.newEmpty();

  let displayError: string = "";
  let disabled: boolean = estudianteMateria.materiaId > 0;
  let materias: Materia[] = [];

  materias = cargaMaterias();

  function cargaMaterias(): Materia[] {
    displayError = "Por favor espere...";
    let materias: Materia[] = [];

    console.log("Cargando lista de materias... ");

    httpClient
      .get("materias")
      .then((response) => {
        displayError = "";
        console.log(response.data);

        // Llena la tabla
        if (!response.data.data) {
          console.log("Empty materias list");
          return;
        }
        let item: Materia;
        for (item of response.data.data) {
          materias.push(item);
        }

        console.log("Lista de materias:");
        console.log(materias);
      })
      .catch((error) => {
        displayError = error;
      });
    return materias;
  }
</script>

<!-- {@debug estudianteMateria} -->

<span class="dispErr">{displayError}</span>
<div class="agrupa">
  <p class="subtitulo">Datos generales</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={10}>
      <Select
        bind:value={estudianteMateria.materiaId}
        label="Seleccione materia"
        {disabled}
        style="width:100%;"
      >
        {#each materias as materia}
          <Option value={materia.id} selected={materia.id === estudianteMateria.materiaId}
            >{materia.nombre}</Option
          >
        {/each}
      </Select>
    </Cell>
  </LayoutGrid>
</div>
<div class="agrupa">
  <p class="subtitulo">Periodo</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={3}>
      <Textfield
        required
        type="date"
        input$maxlength="10"
        style="width: 100%;"
        bind:value={estudianteMateria.inicio}
        label="Inicio"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="date"
        input$maxlength="10"
        style="width: 100%;"
        bind:value={estudianteMateria.fin}
        label="Finalización"
        helperLine$style="width: 100%;"
      />
    </Cell>
  </LayoutGrid>
</div>
<LayoutGrid style="padding:4px 0px;">
  <Cell span={12}>
    <Textfield
      textarea
      type="text"
      input$maxlength="100"
      style="width: 100%;"
      bind:value={estudianteMateria.observaciones}
      label="Observaciones"
      helperLine$style="width: 100%;"
    />
  </Cell>
</LayoutGrid>

<style>
  .subtitulo {
    margin: 0px;
    padding: 5px 0px 0px 0px;
    font-weight: 500;
  }

  .agrupa {
    border: black solid 2px;
    border-radius: 4px;
    padding: 0px 5px 2px 5px;
    margin: 3px 0px;
  }

  .dispErr {
    color: red;
  }
</style>
