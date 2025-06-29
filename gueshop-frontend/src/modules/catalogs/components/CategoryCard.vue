<template>
    <div class="card bg-base-200 w-80 drop-shadow-xl">
        <figure>
            <img :src="props.category.imageUrl ?? notFoundUrl" alt="Joyas" />
        </figure>
        <div class="card-body">
            <h2 class="card-title">{{ props.category.name }}</h2>
            <p> {{ props.category.description }}</p>
            <div class="card-actions justify-end">
                <button class="btn btn-primary">Ver</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import type { Category } from '../types/catalog';

const notFoundUrl = 'https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp';

const props = defineProps<{
    category: Category
}>()

</script>


<style scoped></style>