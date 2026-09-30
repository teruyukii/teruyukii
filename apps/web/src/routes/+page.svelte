<script>
	import Artwork from '$lib/components/Artwork.svelte';
	import Width21 from '$lib/layout/Width21.svelte';
</script>

<Width21>
	<Artwork />
	{#snippet sidebar()}
		<h1>Sidebar</h1>
	{/snippet}
</Width21>
