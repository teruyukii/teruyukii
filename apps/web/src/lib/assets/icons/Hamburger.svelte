<script>
	let { color = 'current color' } = $props();
</script>

<svg fill="none" height="18" viewBox="0 0 18 18" width="18" xmlns="http://www.w3.org/2000/svg">
	<path
		d="M3 3.75H15M3 9H15M3 14.25H15"
		stroke={color}
		stroke-linecap="round"
		stroke-linejoin="round"
		stroke-width="1.5"
	/>
</svg>
