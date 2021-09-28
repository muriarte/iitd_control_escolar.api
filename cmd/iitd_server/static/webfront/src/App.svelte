<script lang="ts">
	import MainPage from "./components/mainpage.svelte";

	export let title: string;
	export let subtitle: string;
	const urlParams = new URLSearchParams(window.location.search);
	const initialModule= urlParams.get("im")
</script>

<MainPage title={title} subtitle={subtitle} initialModule={initialModule} />

<style>
</style>