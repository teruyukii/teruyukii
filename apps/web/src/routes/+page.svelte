<script>
	import Width21 from '$lib/layout/Width21.svelte';
</script>

<Width21>
	<h1>hello</h1>
	{#snippet sidebar()}
		<h1>Sidebar</h1>
	{/snippet}
</Width21>
