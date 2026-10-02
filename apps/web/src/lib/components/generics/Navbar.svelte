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

<nav class="sticky top-0 z-50 flex bg-[#f5f5f5]">
	<div class="flex w-full items-center justify-between px-4 py-4 md:px-8">
		<a class="flex items-center gap-2 font-display text-lg font-bold" href="/"
			><img alt="My App Logo" class="logo h-4 w-4" src={favicon} />teruyuki</a
		>
		<button
			class="md:hidden"
			onclick={() => {
				isOpen = !isOpen;
				log(isOpen ? 'Menu opened' : 'Menu closed');
			}}
		>
			<Hamburger />
		</button>

		<div class="hidden items-center gap-6 md:flex">
			<a class="logo" href="/">stories</a>
			<a class="nav-link" href="/about">works</a>
			<a class="nav-link" href="/contact">me</a>
			<a class="nav-link flex items-center gap-1" href="/contact"
				>github<OutsideLink color="#2f2f2f" /></a
			>
			<a class="nav-link flex items-center gap-1" href="/contact"
				>linkedIn<OutsideLink color="#2f2f2f" /></a
			>
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
