<script lang="ts"></script>

<svelte:head>
	<title>登录</title>
</svelte:head>

<div class="flex flex-col items-center w-full h-full">
	<div class="prose lg:prose-xl">
		<h1>Title</h1>
	</div>
	<div class="flex card my-28 mx-2 md:mx-8 bordered bg-base-200 shadow-xl">
		<div class="md:p-16 card-body">
			<h2 class="card-title">登录管理账户</h2>
			<div class="flex flex-col space-y-2">
				<div class="form-control">
					<input type="text" placeholder="用户名" class="input input-lg" />
				</div>
				<div class="form-control">
					<div class="relative">
						<input
							type="password"
							placeholder="密码"
							class="w-full pr-16 input input-lg"
						/>
						<button class="absolute top-0 right-0 rounded-l-none btn btn-lg btn-ghost btn-primary">go</button>
					</div>
				</div>
			</div>
			<div class="card-actions">
				<button class="btn btn-primary">Get Started</button>
				<button class="btn btn-ghost">More info</button>
			</div>
		</div>
	</div>
</div>
