<script lang="ts">
	import '$lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';
	import OutsideLink from '$lib/assets/icons/OutsideLink.svelte';
	import Hamburger from '$lib/assets/icons/Hamburger.svelte';

	function log(message: string) {
		console.log(message);
	}
	let isOpen = $state(false);
</script>

<nav class="sticky top-0 z-50 flex bg-[#181D1D]">
	<div class="flex w-full items-center justify-between px-4 py-4 md:px-8">
		<a href="/" class="flex items-center gap-2 font-display text-lg font-bold"
			><img src={favicon} alt="My App Logo" class="logo h-4 w-4" />teruyuki</a
		>
		<button
			onclick={() => {
				isOpen = !isOpen;
				log(isOpen ? 'Menu opened' : 'Menu closed');
			}}
			class="md:hidden"
		>
			<Hamburger />
		</button>

		<div class="hidden items-center gap-6 md:flex">
			<a href="/" class="logo">stories</a>
			<a href="/about" class="nav-link">works</a>
			<a href="/contact" class="nav-link">me</a>
			<a href="/contact" class="nav-link flex gap-1">github<OutsideLink /></a>
			<a href="/contact" class="nav-link flex items-center gap-1">linkedIn<OutsideLink /></a>
		</div>
	</div>
</nav>

{#if isOpen}
	<div class="min-h-screen">
		<a href="/" class="nav-link">stories</a>
		<a href="/about" class="nav-link">works</a>
		<a href="/contact" class="nav-link">me</a>
	</div>
{/if}
