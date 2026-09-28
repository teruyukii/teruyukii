<script lang="ts">
	import '$lib/styles/global.css';

	let { children, sidebar } = $props();
</script>

<div class="flex h-fit min-h-dvh w-full items-start p-8">
	<section class="flex-co flex w-full max-w-2/3">{@render children()}</section>
	<div class="flex w-full max-w-1/3 flex-col">
		{@render sidebar()}
	</div>
</div>
