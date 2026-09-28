<script>
	import Breadcrumb from '$lib/components/Breadcrumb.svelte';
	import Width21 from '$lib/layout/Width21.svelte';
</script>

<Width21>
	<Breadcrumb />
	{#snippet sidebar()}
		<h1>Teruyuki Saito</h1>
		<p>
			Coding excites me because it allows me to bring creative ideas to life while maintaining
			structure and organization, which I truly enjoy.
		</p>
		<p>
			Outside of programming, one of my favorite hobbies is playing Karuta, a traditional Japanese
			card game. I’m also an active member of the Bangkok Karuta Club, where I get to engage with
			others who share this interest.
		</p>
		<p>
			Additionally, I have a growing interest in religious studies and philosophy, particularly
			Buddhist philosophy. I plan to include topics related to philosophy on this site, so you can
			get to know me better and see the various aspects of who I am.
		</p>
		<p>
			I’m always eager to explore new opportunities for learning and personal growth, so feel free
			to reach out to me anytime!
		</p>
	{/snippet}
</Width21>
