<script lang="ts">
	import { Label, Input, Button, A, Helper } from 'flowbite-svelte';
	import { EyeOutline, EyeSlashOutline } from 'flowbite-svelte-icons';
	let show: boolean = false;
	let show2: boolean = false;

	let username: string = '';
	let email: string = '';
	let password: string = '';
	let passwordverif: string = '';
</script>

<svelte:head>
	<title>Create Account</title>
</svelte:head>
<form action="">
	<h2 class="mb-8 text-2xl font-bold">Create Account</h2>
	<div class="mb-6">
		<Label for="username" class="text-md mb-2 block">Username</Label>
		<Input
			type="text"
			id="username"
			placeholder="Your Awesome Username"
			class="text-md border-2"
			name="username"
			autocomplete="username"
			bind:value={username}
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="email" class="text-md mb-2 block">Email</Label>
		<Input
			type="email"
			id="email"
			placeholder="john.doe@company.com"
			class="text-md border-2"
			name="email"
			autocomplete="email"
			bind:value={email}
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="password" class="text-md mb-2 block">Password</Label>
		<Input
			type={show ? 'text' : 'password'}
			id="password"
			placeholder="•••••••••"
			class="text-md border-2"
			bind:value={password}
			name="password"
			required
		>
			<button slot="left" on:click={() => (show = !show)} class="pointer-events-auto">
				{#if !show}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</Input>
	</div>
	<div class="mb-6">
		<Label
			for="password-verification"
			class="text-md mb-2 block"
			color={password === passwordverif ? 'gray' : 'red'}>Password Verification</Label
		>
		<Input
			type={show ? 'text' : 'password'}
			id="password-verification"
			placeholder="•••••••••"
			class="text-md border-2"
			bind:value={passwordverif}
			color={password === passwordverif ? 'base' : 'red'}
			required
		>
			<button slot="left" on:click={() => (show2 = !show2)} class="pointer-events-auto">
				{#if !show2}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</Input>
		{#if password !== passwordverif}
			<Helper color="red">Password harus sama</Helper>
		{/if}
	</div>
	<p class="mb-4">
		Sudah punya akun? <A
			href="/auth/login"
			class="text-primary-700 hover:underline dark:text-primary-600">login sekarang</A
		>
	</p>
	<Button type="submit" disabled={password !== passwordverif || password == '' || email == ''}
		>Sign In</Button
	>
</form>
