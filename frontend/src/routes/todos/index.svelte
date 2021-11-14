<script context="module" lang="ts">
	import { enhance } from '$lib/form';
	import type { Load } from '@sveltejs/kit';

	// see https://kit.svelte.dev/docs#loading
	export const load: Load = async ({ fetch }) => {
		const res = await fetch('/todos.json');

		if (res.ok) {
			const todos = await res.json();

			return {
				props: { todos }
			};
		}

		const { message } = await res.json();

		return {
			error: new Error(message)
		};
	};
</script>

<script lang="ts">
	import { scale } from 'svelte/transition';
	import { flip } from 'svelte/animate';

	type Todo = {
		uid: string;
		created_at: Date;
		text: string;
		done: boolean;
		pending_delete: boolean;
	};

	export let todos: Todo[];

	async function patch(res: Response) {
		const todo = await res.json();

		todos = todos.map((t) => {
			if (t.uid === todo.uid) return todo;
			return t;
		});
	}
</script>

<svelte:head>
	<title>Todos</title>
</svelte:head>

<div class="todos">
	<div
		class="flex items-center w-full px-4 py-10 bg-cover card bg-base-200"
		style="background-image: url(&quot;https://picsum.photos/id/314/1000/300&quot;);"
	>
		<div class="card glass lg:card-side text-neutral-content">
			<figure class="p-6">
				<img src="https://picsum.photos/id/1005/300/200" class="rounded-lg shadow-lg" />
			</figure>
			<div class="max-w-md card-body">
				<h2 class="card-title">Glass</h2>
				<p>
					Rerum reiciendis beatae tenetur excepturi aut pariatur est eos. Sit sit necessitatibus
					veritatis sed molestiae voluptates incidunt iure sapiente.
				</p>
				<div class="card-actions">
					<button class="btn glass rounded-full">Get Started</button>
				</div>
			</div>
		</div>
	</div>
</div>
