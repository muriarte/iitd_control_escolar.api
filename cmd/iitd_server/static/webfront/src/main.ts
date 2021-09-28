import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		title: 'IITD',
		subtitle: 'Control escolar'
	}
});

export default app;