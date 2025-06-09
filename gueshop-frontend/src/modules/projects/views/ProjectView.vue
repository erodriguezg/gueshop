<template>
    <div class="overflow-x-auto w-full">
        <table class="table">
            <!-- head -->
            <thead>
                <tr>
                    <th></th>
                    <th>Proyecto</th>
                    <th>Tareas</th>
                    <th>Avance</th>
                </tr>
            </thead>
            <tbody>
                <!-- row 2 -->
                <tr class="hover:bg-base-300">
                    <th>2</th>
                    <td>Hart Hagerty</td>
                    <td>Desktop Support Technician</td>
                    <td>Purple</td>
                </tr>
            </tbody>
        </table>
    </div>

    <fab-button position="bottom-right" />

</template>

<script setup lang="ts">
import FabButton from '@/modules/common/components/FabButton.vue';

</script>