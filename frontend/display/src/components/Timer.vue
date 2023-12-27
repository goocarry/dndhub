<template>
    <div class="mx-4">
        <p class="text-h4 font-weight-black text-white">До старта марафона осталось:</p>
    </div>
    <div class="mx-4 text-white">
        {{ displayDays }}
        {{ displayHours }}
        {{ displayMinutes }}
        {{ displaySeconds }}
    </div>
</template>

<script>
export default {
    data: () => ({
        displayDays: 0,
        displayHours: 0,
        displayMinutes: 0,
        displaySeconds: 0,
    }),
}
</script>