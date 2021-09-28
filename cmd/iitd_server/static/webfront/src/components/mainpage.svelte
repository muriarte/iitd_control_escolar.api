<script lang="ts">
  import Drawer, {
    AppContent,
    Content,
    Header,
    Title,
    Subtitle,
    Scrim,
  } from "@smui/drawer/styled";
  import Button, { Label } from "@smui/button/styled";
  import IconButton from "@smui/icon-button/styled";
  import List, {
    Item,
    Text,
    Graphic,
    Separator,
    Subheader,
  } from "@smui/list/styled";
  import H6 from "@smui/common/H6.svelte";

  // import Splash from "./splash.svelte";
  // import BuscaFactura from "./busca-factura.svelte";

  const modCatEstudiantes: string = "CatalogoEstudiantes";
  const modSplash: string = "Splash";

  export let title: string = "";
  export let subtitle: string = "";
  export let initialModule: string = modSplash;

  let open = false;

  let active = initialModule;

  function setActive(value: string) {
    active = value;
    open = false;
  }
</script>

<!-- <div class="drawer-container"> -->
<!-- Don't include fixed={false} if this is a page wide drawer.
            It adds a style for absolute positioning. -->
<!-- <Drawer variant="modal" fixed={false} bind:open> -->
<Drawer variant="modal" bind:open>
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
        <Graphic class="material-icons" aria-hidden="true">search</Graphic>
        <Text>Catálogo de estudiantes</Text>
      </Item>

      <Separator />
      <Subheader component={H6}>Labels</Subheader>
      <Item
        href="javascript:void(0)"
        on:click={() => setActive("Logout")}
        activated={active === "Logout"}
      >
        <Graphic class="material-icons" aria-hidden="true">exit</Graphic>
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
    <IconButton
      class="material-icons"
      on:click={() => (open = !open)}
      ripple={false}>menu</IconButton
    >
    <!-- <br /> -->
    <!-- <pre class="status">Active: {active}</pre> -->
    {#if active == modSplash}
      <!-- <Splash /> -->
    {/if}
    {#if active == modCatEstudiantes}
      <!-- <BuscaFactura /> -->
    {/if}
  </main>
</AppContent>

<!-- </div> -->
<style>
  /* These classes are only needed because the
        drawer is in a container on the page. */
  /* .drawer-container {
      position: relative;
      display: flex;
      height: 350px;
      max-width: 600px;
      border: 1px solid
        var(--mdc-theme-text-hint-on-background, rgba(0, 0, 0, 0.1));
      overflow: hidden;
      z-index: 0;
    } */

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
</style>
