<script lang="ts">
	import PasswordInput from '$lib/input/PasswordInput.svelte';
	import FormInput from '$lib/input/FormInput.svelte';

	let username = '';
	let password = '';

	function handleSubmit() {
		console.log(username);
		console.log(password);
	}
</script>

<svelte:head>
	<title>登录</title>
</svelte:head>

<div class="flex items-center justify-center w-full h-full">
	<div class="flex card my-28 mx-2 md:mx-8 bordered shadow-2xl bg-base-200">
		<div class="md:p-12 card-body">
			<div class="card-title prose lg:prose-xl pb-4">
				<h2>登录Shifter</h2>
			</div>
			<div class="flex flex-col space-y-2">
				<FormInput bind:value={username} placeholder="用户名" size="lg" bordered />
				<PasswordInput bind:value={password} placeholder="密码" size="lg" bordered withShowButton />
			</div>
			<div class="card-actions justify-center">
				<!-- <button class="btn btn-ghost">找回密码</button> -->
				<button class="btn btn-primary w-1/3" on:click={handleSubmit}>登录</button>
			</div>
		</div>
	</div>
</div>
