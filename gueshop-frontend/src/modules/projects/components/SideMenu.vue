<template>
    <aside class="bg-base-200 w-72 min-h-screen">
        <p class="text-sm text-gray-500 mx-4">No hay proyectos</p>

        <!-- Menu -->
        <ul class="menu w-full">
            <li><a>Item 1</a></li>
            <li>
                <details open>
                    <summary>Parent</summary>
                    <ul>
                        <li><a>Submenu 1</a></li>
                        <li><a>Submenu 2</a></li>
                        <li>
                            <details open>
                                <summary>Parent</summary>
                                <ul>
                                    <li><a>Submenu 1</a></li>
                                    <li><a>Submenu 2</a></li>
                                </ul>
                            </details>
                        </li>
                    </ul>
                </details>
            </li>
            <li><a>Item 3</a></li>
        </ul>
    </aside>
</template>