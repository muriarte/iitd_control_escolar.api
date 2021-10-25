<script lang="ts">
  import Drawer, { AppContent, Content,Header, Title, Subtitle, Scrim } from "@smui/drawer/styled";
  import Button, { Label } from "@smui/button/styled";
  import IconButton from "@smui/icon-button/styled";
  import List, { Item, Text, Graphic, Separator, Subheader } from "@smui/list/styled";
  import H6 from "@smui/common/H6.svelte";

  import Splash from "./splash.svelte";
  import CatalogoEstudiantes from "./catalogoestudiantes.svelte";
  import CatalogoMaestros from "./catalogomaestros.svelte";
  import CatalogoMaterias from "./catalogomaterias.svelte";

  const modCatEstudiantes: string = "CatalogoEstudiantes";
  const modCatMaestros: string = "CatalogoMaestros";
  const modCatMaterias: string = "CatalogoMaterias";
  const modSplash: string = "Splash";

  export let title: string = "";
  export let subtitle: string = "";
  export let initialModule: string = modSplash;

  let drawerIsOpen = false;

  let active = initialModule;
  if (!active) {
    active = modSplash;
  }

  function setActive(value: string) {
    active = value;
    drawerIsOpen = false;
  }
  
</script>

<!-- <div class="drawer-container"> -->
<!-- Don't include fixed={false} if this is a page wide drawer.
            It adds a style for absolute positioning. -->
<!-- <Drawer variant="modal" fixed={false} bind:open> -->
<Drawer variant="modal" bind:open={drawerIsOpen}>
<Header>
  <Title>{title}</Title>
  <Subtitle>{subtitle}</Subtitle>
</Header>
<Content>
    <List>
      <Item
        href="javascript:void(0)"
        on:click={() => setActive(modCatEstudiantes)}
        activated={active === modCatEstudiantes}
      >
        <Graphic class="material-icons" aria-hidden="true">school</Graphic>
        <Text>Catálogo de estudiantes</Text>
      </Item>

      <Item
        href="javascript:void(0)"
        on:click={() => setActive(modCatMaestros)}
        activated={active === modCatMaestros}
      >
        <Graphic class="material-icons" aria-hidden="true">record_voice_over</Graphic>
        <Text>Catálogo de maestros</Text>
      </Item>

      <Item
        href="javascript:void(0)"
        on:click={() => setActive(modCatMaterias)}
        activated={active === modCatMaterias}
      >
        <Graphic class="material-icons" aria-hidden="true">book</Graphic>
        <Text>Catálogo de materias</Text>
      </Item>

      <Separator />
      <Subheader component={H6}>Labels</Subheader>
      <Item href="javascript:void(0)" on:click={() => setActive("Logout")} activated={active === "Logout"}>
        <Graphic class="material-icons" aria-hidden="true">logout</Graphic>
        <Text>Logout</Text>
      </Item>
    </List>
  </Content>
</Drawer>

<!-- Don't include fixed={false} if this is a page wide drawer.
            It adds a style for absolute positioning. -->
<!-- <Scrim fixed={false} />  -->
<Scrim />
<AppContent class="app-content">
  <main class="main-content">
    <IconButton class="material-icons" on:click={() => (drawerIsOpen = !drawerIsOpen)} ripple={false}>menu</IconButton>
    <h3>{title} - {subtitle}</h3><br>
    <!-- active:{active} -->
    {#if active == modSplash}
      <Splash />
    {/if}
    {#if active == modCatEstudiantes}
      <CatalogoEstudiantes />
    {/if}
    {#if active == modCatMaestros}
      <CatalogoMaestros />
    {/if}
    {#if active == modCatMaterias}
      <CatalogoMaterias />
    {/if}
  </main>
</AppContent>

<!-- </div> -->
<style>
  * :global(.app-content) {
    flex: auto;
    overflow: auto;
    position: relative;
    flex-grow: 1;
  }

  .main-content {
    overflow: auto;
    padding: 5px;
    height: 100%;
    box-sizing: border-box;
  }
  h3 {
    display:inline;
  }
</style>
