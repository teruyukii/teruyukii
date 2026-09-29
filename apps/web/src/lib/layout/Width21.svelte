<script lang="ts">
	import Footer from '$lib/components/generics/Footer.svelte';
	import '$lib/styles/global.css';

	let { children, sidebar } = $props();
</script>

<div class="flex h-[calc(100dvh-60px)] w-full flex-col items-start gap-4 p-8 md:flex-row">
	<section class="flex w-full flex-col md:max-w-2/3">{@render children()}</section>
	<section class="flex h-full w-full flex-col justify-between md:max-w-1/3">
		<div>
			{@render sidebar()}
		</div>
		<Footer />
	</section>
</div>
