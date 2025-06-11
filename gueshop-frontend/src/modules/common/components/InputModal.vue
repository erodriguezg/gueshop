<template>
    <dialog id="my_modal_1" class="modal" :open="true">
        <div class="modal-box">
            <h3 class="text-lg font-bold">Hello!</h3>
            <p class="py-4">Press ESC key or click the button below to close</p>

            <div class="modal-action flex flex-col">
                <form method="dialog" @submit.prevent="submitValue">
                    <input type="text" placeholder="Nombre del proyecto"
                        class="input input-bordered input-primary w-full flex-1" v-model="inputValue" />

                    <!-- if there is a button in form, it will close the modal -->
                    <div class="flex justify-end mt-5">
                        <button class="btn mr-4">Close</button>
                        <button type="submit" class="btn btn-primary">Aceptar</button>
                    </div>

                </form>
            </div>

        </div>
    </dialog>

    <div class="modal-backdrop fixed top-0 left-0 z-10 bg-black opacity-40 w-screen h-screen"></div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

interface Props {
    open: boolean;
}

const emits = defineEmits<{
    close: [void];
    value: [text: string];
}>()

const inputValue = ref('');

const submitValue = () => {
    console.log({
        value: inputValue.value
    });

    if (!inputValue.value) {
        // foco en el elemento
        return;
    }
    emits('value', inputValue.value.trim())
    emits('close');

    inputValue.value = '';
}



</script>