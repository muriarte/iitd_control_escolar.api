<script lang="ts">
  import { Router, Link, Route, navigate } from "svelte-navigator";
  import Drawer, { AppContent, Content, Header, Title, Subtitle, Scrim } from "@smui/drawer/styled";
  import IconButton from "@smui/icon-button/styled";
  import List, { Item, Text, Graphic, Separator, Subheader } from "@smui/list/styled";
  import H6 from "@smui/common/elements/H6.svelte";

  import Splash from "./splash.svelte";
  import CatalogoEstudiantes from "./catalogoestudiantes.svelte";
  import CatalogoMaestros from "./catalogomaestros.svelte";
  import CatalogoMaterias from "./catalogomaterias.svelte";
  import EstudianteMaterias from "./estudiantematerias.svelte";
  import EstudianteObs from "./estudianteobs.svelte";

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
    if (value === modCatEstudiantes) {
      navigate("/estudiantes");
    }
    if (value === modCatMaestros) {
      navigate("/maestros");
    }
    if (value === modCatMaterias) {
      navigate("/materias");
    }
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
<!-- Scrim crea una sombra que nubla el fondo de la pantalla para que resalte el drawer al abrirlo. -->
<!-- Don't include fixed={false} if this is a page wide drawer.
            It adds a style for absolute positioning. -->
<!-- <Scrim fixed={false} />  -->
<Scrim />
<AppContent class="app-content">
  <main class="main-content">
    <IconButton class="material-icons" on:click={() => (drawerIsOpen = !drawerIsOpen)} ripple={false}
      >menu</IconButton
    >
    <h3>{title} - {subtitle}</h3>
    <br />
    <Router>
      <nav>
        <Link to="/">Home</Link>
        <Link to="estudiantes">Estudiantes</Link>
        <Link to="maestros">Maestros</Link>
        <Link to="materias">Materias</Link>
      </nav>
      <div>
        <Route path="/">
          <Splash />
        </Route>
        <Route path="estudiantes">
          <CatalogoEstudiantes />
        </Route>
        <Route path="maestros">
          <CatalogoMaestros />
        </Route>
        <Route path="materias">
          <CatalogoMaterias />
        </Route>
        <Route path="estudiantematerias/:estudianteId" let:params>
          <EstudianteMaterias studentId={parseInt(params.estudianteId)} />
        </Route>
        <Route path="estudianteobs/:estudianteId" let:params>
          <EstudianteObs studentId={parseInt(params.estudianteId)} />
        </Route>
      </div>
    </Router>
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
    display: inline;
  }

  * :global(nav a) {
    font-family: roboto;
    padding-right: 0.5em;
  }
</style>
