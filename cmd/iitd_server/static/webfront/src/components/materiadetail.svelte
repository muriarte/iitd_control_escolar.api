<script lang="ts">
  import Textfield from "@smui/textfield/styled";
  import LayoutGrid, { Cell } from "@smui/layout-grid/styled";
  import Radio from "@smui/radio/styled";
  import FormField from "@smui/form-field/styled";
  import { Label } from "@smui/common/styled";
  import Checkbox from "@smui/checkbox/styled";
  import type { Materia } from "../models/materia";

  export let materia: Materia = {
    id: 0,
    nombre: "",
    observaciones: "",
    activo: "N",
  };
  $: activo = materia.activo === "S";
</script>

<div class="agrupa">
  <p class="subtitulo">Datos generales</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={12}>
      <Textfield
        required
        type="text"
        input$maxlength="100"
        style="width: 100%;"
        bind:value={materia.nombre}
        label="Nombre"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={12}>
      <Textfield
        textarea
        type="text"
        input$maxlength="100"
        style="width: 100%;"
        bind:value={materia.observaciones}
        label="Observaciones"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <FormField>
        <Checkbox
          checked={activo}
          on:change={(e) => {
            materia.activo = e?.srcElement?.checked ? "S" : "N";
            console.log("Activo changed to " + materia.activo + "," + e?.srcElement?.checked);
            console.log(e);
          }}
          touch
        />
        <span slot="label">Activo.</span>
      </FormField>
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
</style>
