<template>
    <div v-if="isSelected">
        <slot></slot>
    </div>
</template>

<script>
export default {
    name:"tab",
    props:{
        isSelected: {
            type : Boolean
        }
    }
}
</script>