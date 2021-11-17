<script lang="ts">
	import EyeOff from '$lib/icon/EyeOff.svelte';
	import Eye from '$lib/icon/Eye.svelte';

	export let value: string;
	export let placeholder = '';
	export let size = '';
    export let state = '';
	export let bordered: boolean = false;
    export let withShowButton: boolean = false; 
	let showPassword = false;

    $: console.log(state)
	$: type = showPassword ? 'text' : 'password';

	const handleShowPassword = () => (showPassword = !!!showPassword);
	const handleInput = (e) => (value = e.target.value)
</script>

<div class="form-control">
	<div class="relative">
		<input
			{type}
			{value}
			{placeholder}
			class="w-full pr-16 input"
			class:input-bordered={bordered}
			class:input-xs={size === 'xs'}
			class:input-sm={size === 'sm'}
			class:input-lg={size === 'lg'}
            class:input-info={state === 'info'}
            class:input-success={state === 'success'}
            class:input-warning={state === 'warning'}
            class:input-error={state === 'error'}
			on:input={handleInput}
		/>
        {#if withShowButton}
		<button
			class="absolute top-0 right-0 rounded-l-none btn btn-lg btn-ghost btn-primary"
			on:click={handleShowPassword}
		>
			{#if showPassword}
				<EyeOff />
			{:else}
				<Eye />
			{/if}
		</button>
        {/if}
	</div>
</div>
