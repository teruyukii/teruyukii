<script>
	let { color = 'current color' } = $props();
</script>

<svg fill="none" height="12" viewBox="0 0 12 12" width="12" xmlns="http://www.w3.org/2000/svg">
	<path
		d="M8.5 8.5V3.5H3.5M8.5 3.5L3.5 8.5"
		stroke={color}
		stroke-linecap="round"
		stroke-linejoin="round"
	/>
</svg>
