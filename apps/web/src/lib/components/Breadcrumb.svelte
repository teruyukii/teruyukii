<script lang="ts">
	import { page } from '$app/stores';
</script>

<div class="breadcrumb flex items-center gap-3 pb-6 font-display font-semibold">
	<a
		href="/"
		class="hover:opacity-100 {($page.url.pathname as string) == '/' ? 'opacity-100' : 'opacity-50'}"
		>Contribution</a
	>
	<a
		href="/about"
		class="hover:opacity-100 {($page.url.pathname as string) == '/about'
			? 'opacity-100'
			: 'opacity-50'}">Centificates</a
	>
	<a
		href="/about/team"
		class="hover:opacity-100 {($page.url.pathname as string) == '/about/team'
			? 'opacity-100'
			: 'opacity-50'}">Team</a
	>
</div>
