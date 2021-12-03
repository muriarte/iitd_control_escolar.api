<script lang="ts">
  import Textfield from "@smui/textfield/styled";
  import LayoutGrid, { Cell } from "@smui/layout-grid/styled";
  import Radio from "@smui/radio/styled";
  import FormField from "@smui/form-field/styled";
  import { Label } from "@smui/common/styled";
  import Checkbox from "@smui/checkbox/styled";
  import Select, { Option } from "@smui/select/styled";
  import { EstudianteObs } from "../models/estudianteobs";
  import type { Materia } from "../models/materia";
  import axios from "axios/dist/axios";

  let urlpathprefix: string = location.protocol + "//" + location.host + "/v1/";
  const httpClient = axios.create({
    timeout: 180000,
    baseURL: urlpathprefix,
  });

  export let estudianteObs: EstudianteObs = EstudianteObs.newEmpty();

  let displayError: string = "";
</script>

<!-- {@debug estudianteObs} -->

<span class="dispErr">{displayError}</span>
<div class="agrupa">
  <p class="subtitulo">Datos generales</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={10}>
      <Textfield required type="date" input$maxlength={10} label="Fecha" bind:value={estudianteObs.fecha} />
    </Cell>
    <Cell span={10}>
      <Textfield
        required
        textarea
        type="text"
        input$maxlength={1000}
        style="width: 100%;"
        bind:value={estudianteObs.observacion}
        label="Observaciones"
        helperLine$style="width: 100%;"
      />
    </Cell>
  </LayoutGrid>
</div>

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
