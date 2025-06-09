<template>

    <button :class="['btn btn-circle btn-secondary fixed', position]">
        <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
            <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
            <g id="SVGRepo_iconCarrier">
                <path d="M15 12L12 12M12 12L9 12M12 12L12 9M12 12L12 15" stroke="#1C274C" stroke-width="1.5"
                    stroke-linecap="round"></path>
                <path
                    d="M7 3.33782C8.47087 2.48697 10.1786 2 12 2C17.5228 2 22 6.47715 22 12C22 17.5228 17.5228 22 12 22C6.47715 22 2 17.5228 2 12C2 10.1786 2.48697 8.47087 3.33782 7"
                    stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"></path>
            </g>
        </svg>
    </button>

</template>

<script lang="ts" setup>

interface Props {
    position?: 'top-left' | 'top-right' | 'bottom-left' | 'bottom-right';
}

withDefaults(defineProps<Props>(), {
    position: 'bottom-right'
});

</script>


<style scoped>
@reference '@/assets/main.css'

.top-left {
    @apply top-10 left-10;
}

.top-right {
    @apply top-10 right-10;
}

.bottom-left {
    @apply bottom-10 left-10;
}

.bottom-right {
    @apply bottom-10 right-10;
}
</style>