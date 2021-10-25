<script lang="ts">
  import Textfield from "@smui/textfield/styled";
  import LayoutGrid, { Cell } from "@smui/layout-grid/styled";
  import Radio from "@smui/radio/styled";
  import FormField from "@smui/form-field/styled";
  import { Label } from "@smui/common/styled";
  import Checkbox from "@smui/checkbox/styled";
  import type { Maestro } from "../models/maestro";

  export let maestro: Maestro = {
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
    telCelular: "",
    telCasa: "",
    email: "",
    fechaInicio: "2000-01-01",
    observaciones: "",
    activo: "N",
  };
  $: activo = maestro.activo === "S";
</script>

<div class="agrupa">
  <p class="subtitulo">Datos generales</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="100"
        style="width: 100%;"
        bind:value={maestro.nombres}
        label="Nombres"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="100"
        style="width: 100%;"
        bind:value={maestro.apellidos}
        label="Apellidos"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="date"
        input$maxlength="10"
        style="width: 100%;"
        bind:value={maestro.nacimiento}
        label="Nacimiento"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Label>Sexo:</Label>
      <FormField>
        <Radio bind:group={maestro.sexo} value="M" name="sexogrp" />
        <span slot="label"> Masculino </span>
      </FormField>
      <FormField>
        <Radio bind:group={maestro.sexo} value="F" name="sexogrp" />
        <span slot="label"> Femenino </span>
      </FormField>
    </Cell>
  </LayoutGrid>
</div>
<div class="agrupa">
  <p class="subtitulo">Domicilio</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="100"
        style="width: 100%;"
        bind:value={maestro.calle}
        label="Calle"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.numeroExt}
        label="# exterior"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.numeroInt}
        label="# interior"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.colonia}
        label="colonia"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.municipio}
        label="municipìo"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.estado}
        label="estado"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.pais}
        label="pais"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="text"
        input$maxlength="6"
        style="width: 100%;"
        bind:value={maestro.cp}
        label="Código Postal"
        helperLine$style="width: 100%;"
      />
    </Cell>
  </LayoutGrid>
</div>
<div class="agrupa">
  <p class="subtitulo">Localización</p>
  <LayoutGrid style="padding:0px 0px;">
    <Cell span={3}>
      <Textfield
        type="tel"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.telCelular}
        label="Teléfono celular"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        type="tel"
        input$maxlength="50"
        style="width: 100%;"
        bind:value={maestro.telCasa}
        label="Teléfono casa"
        helperLine$style="width: 100%;"
      />
    </Cell>

    <Cell span={3}>
      <Textfield
        required
        type="email"
        input$maxlength="255"
        style="width: 100%;"
        bind:value={maestro.email}
        label="email"
        helperLine$style="width: 100%;"
      />
    </Cell>
  </LayoutGrid>
</div>
<LayoutGrid style="padding:0px 0px;">
  <Cell span={3}>
    <Textfield
      required
      type="date"
      input$maxlength="10"
      style="width: 100%;"
      bind:value={maestro.fechaInicio}
      label="Fecha de inicio"
      helperLine$style="width: 100%;"
    />
  </Cell>

  <Cell span={3}>
    <Textfield
      required
      input$rows="5"
      input$resizable={true}
      type="text"
      style="width: 100%;"
      bind:value={maestro.observaciones}
      label="Observaciones"
      helperLine$style="width: 100%;"
    />
  </Cell>

  <Cell span={3}>
    <FormField>
      <Checkbox
        checked={activo}
        on:change={(e) => {
          maestro.activo = e?.srcElement?.checked ? "S" : "N";
          console.log("Activo changed to " + maestro.activo + "," + e?.srcElement?.checked);
          console.log(e);
        }}
        touch
      />
      <span slot="label">Activo.</span>
    </FormField>
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
</style>
