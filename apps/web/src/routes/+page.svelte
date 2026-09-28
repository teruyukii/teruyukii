<script>
	import Breadcrumb from '$lib/components/Breadcrumb.svelte';
	import Width21 from '$lib/layout/width21.svelte';
	let { data } = $props();
</script>

<Width21>
	<h1>hello</h1>
	{#snippet sidebar()}
		<h1>Sidebar</h1>
	{/snippet}
</Width21>
