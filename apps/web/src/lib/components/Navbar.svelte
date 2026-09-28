<script lang="ts">
	import '$lib/styles/global.css';
	import favicon from '$lib/assets/favicon.svg';
	import OutsideLink from '$lib/assets/icons/OutsideLink.svelte';
</script>

<nav class="flex w-full items-center justify-between px-8 py-4">
	<a href="/" class="flex items-center gap-2 font-display text-lg font-bold"
		><img src={favicon} alt="My App Logo" class="logo h-4 w-4" />teruyuki</a
	>
	<div class="flex items-center gap-6">
		<a href="/" class="logo">stories</a>
		<a href="/about" class="nav-link">works</a>
		<a href="/contact" class="nav-link">me</a>
		<a href="/contact" class="nav-link flex gap-1">github<OutsideLink /></a>
		<a href="/contact" class="nav-link flex items-center gap-1">linkedIn<OutsideLink /></a>
	</div>
</nav>
