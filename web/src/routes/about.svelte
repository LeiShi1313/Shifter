<script context="module">
	import { browser, dev } from '$app/env';

	// we don't need any JS on this page, though we'll load
	// it in dev so that we get hot module replacement...
	export const hydrate = dev;

	// ...but if the client-side router is already loaded
	// (i.e. we came here from elsewhere in the app), use it
	export const router = browser;

	// since there's no dynamic data here, we can prerender
	// it so that it gets served as a static asset in prod
	export const prerender = true;
</script>

<svelte:head>
	<title>About</title>
</svelte:head>

<div class="content flex flex-col mt-30">

	<h1>About this app</h1>

	<button class="btn btn-active" role="button" aria-pressed="true">neutral</button> 
<button class="btn btn-primary btn-active" role="button" aria-pressed="true">primary</button> 
<button class="btn btn-secondary btn-active" role="button" aria-pressed="true">secondary</button> 
<button class="btn btn-accent btn-active" role="button" aria-pressed="true">accent</button> 
<button class="btn btn-link btn-active" role="button" aria-pressed="true">link</button>

<div class="rounded-lg shadow bg-base-200 drawer h-52">
	<input id="my-drawer" type="checkbox" class="drawer-toggle"> 
	<div class="flex flex-col items-center justify-center drawer-content">
	  <label for="my-drawer" class="btn btn-primary drawer-button">open menu</label>
	</div> 
	<div class="drawer-side">
	  <label for="my-drawer" class="drawer-overlay"></label> 
	  <ul class="menu p-4 overflow-y-auto w-80 bg-base-100 text-base-content">
		<li>
		  <a>Menu Item</a>
		</li> 
		<li>
		  <a>Menu Item</a>
		</li>
	  </ul>
	</div>
  </div>

	<p>
		This is a <a href="https://kit.svelte.dev">SvelteKit</a> app. You can make your own by typing the
		following into your command line and following the prompts:
	</p>

	<!-- TODO lose the @next! -->
	<pre>npm init svelte@next</pre>

	<p>
		The page you're looking at is purely static HTML, with no client-side interactivity needed.
		Because of that, we don't need to load any JavaScript. Try viewing the page's source, or opening
		the devtools network panel and reloading.
	</p>

	<p>
		The <a href="/todos">TODOs</a> page illustrates SvelteKit's data loading and form handling. Try using
		it with JavaScript disabled!
	</p>
</div>

<style>
	.content {
		width: 100%;
		max-width: var(--column-width);
		margin: var(--column-margin-top) auto 0 auto;
	}
</style>
