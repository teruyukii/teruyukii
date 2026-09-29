<script lang="ts">
	import { onMount } from 'svelte';
	import Width21 from '$lib/layout/Width21.svelte';

	const isDesktop = () => typeof window !== 'undefined' && window.innerWidth >= 768;
	let activeTab = $state(isDesktop() ? 'contribution' : 'aboutme');

	function openTab(tabName: string) {
		activeTab = tabName;
	}

	onMount(() => {
		const handleResize = () => {
			if (isDesktop()) {
				if (activeTab === 'aboutme') {
					activeTab = 'contribution';
				}
			} else if (activeTab === 'contribution' && !window.matchMedia('(min-width: 768px)').matches) {
				activeTab = 'aboutme';
			}
		};

		handleResize();
		window.addEventListener('resize', handleResize);

		return () => {
			window.removeEventListener('resize', handleResize);
		};
	});
</script>

<Width21>
	<div class="breadcrumb flex items-center gap-3 pb-6 font-display font-semibold">
		<button
			class="tablinks hover:opacity-100 md:hidden"
			class:opacity-100={activeTab === 'aboutme'}
			class:opacity-50={activeTab !== 'aboutme'}
			onclick={() => openTab('aboutme')}
		>
			About me
		</button>
		<button
			class="tablinks hover:opacity-100"
			class:opacity-100={activeTab === 'contribution'}
			class:opacity-50={activeTab !== 'contribution'}
			onclick={() => openTab('contribution')}
		>
			Contribution
		</button>
		<button
			class="tablinks hover:opacity-100"
			class:opacity-100={activeTab === 'certificates'}
			class:opacity-50={activeTab !== 'certificates'}
			onclick={() => openTab('certificates')}
		>
			Certificates
		</button>
	</div>

	{#if activeTab === 'aboutme'}
		<div id="aboutme" class="tabcontent">
			<h1>Teruyuki Saito</h1>
			<p>
				I have contributed to several open-source projects, including SvelteKit, where I have
				submitted bug fixes and feature enhancements. I also maintain a few personal projects on
				GitHub, which are available for public use and collaboration.
			</p>
			<p>
				In addition to coding, I actively participate in online forums and communities related to
				web development and programming. I enjoy sharing my knowledge and helping others solve
				problems they encounter in their projects.
			</p>
		</div>
	{/if}

	{#if activeTab === 'contribution'}
		<div id="contribution" class="tabcontent">
			<h1>Contribution</h1>
			<p>
				I have contributed to several open-source projects, including SvelteKit, where I have
				submitted bug fixes and feature enhancements. I also maintain a few personal projects on
				GitHub, which are available for public use and collaboration.
			</p>
			<p>
				In addition to coding, I actively participate in online forums and communities related to
				web development and programming. I enjoy sharing my knowledge and helping others solve
				problems they encounter in their projects.
			</p>
		</div>
	{/if}

	{#if activeTab === 'certificates'}
		<div id="certificates" class="tabcontent">
			<h1>Certificates</h1>
			<p>
				I have contributed to several open-source projects, including SvelteKit, where I have
				submitted bug fixes and feature enhancements. I also maintain a few personal projects on
				GitHub, which are available for public use and collaboration.
			</p>
			<p>
				In addition to coding, I actively participate in online forums and communities related to
				web development and programming. I enjoy sharing my knowledge and helping others solve
				problems they encounter in their projects.
			</p>
		</div>
	{/if}
	{#snippet sidebar()}
		<div class="hidden md:block">
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
		</div>
	{/snippet}
</Width21>
