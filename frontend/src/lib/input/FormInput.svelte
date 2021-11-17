<script lang="ts">
	export let value;
	export let placeholder = '';
	export let size = '';
    export let bordered: boolean = false;
</script>

<div class="form-control">
	<input
		type="text"
		bind:value
		{placeholder}
		class="input"
        class:input-bordered={bordered}
		class:input-xs={size === 'xs'}
		class:input-sm={size === 'sm'}
		class:input-lg={size === 'lg'}
	/>
</div>
