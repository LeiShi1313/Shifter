<script context="module" lang="ts">
    export function load({ error, status }) {
        return {
            props: { error, status }
        }
    }
</script>

<script lang="ts">
    import { dev } from "$app/env"

    export let error: { message: any; stack: any; }, status: any
</script>

<svelte:head>
    <title>{status}</title>
</svelte:head>

<div class="">
    <h1 class="text-md font-bold">{status}</h1>

    <p>{error.message}</p>

    {#if dev && error.stack}
        <pre>{error.stack}</pre>
    {/if}
</div>