<template>
  <transition :name="slideDirection" @beforeEnter="beforeEnter" @leave="leave">
    <div
      v-show="isActived"
      :style="{ 'transition-duration': `${transition.duration}ms` }"
      class="tab-item"
    >
      <slot />
    </div>
  </transition>
</template>

<script>
export default {
  name: "TabItem",

  props: {
    name: {
      type: String,
      default: "Tab Item",
    },

    disabled: Boolean,
  },

  data: () => ({
    activeModel: "",
    model: Math.random()
      .toString("16")
      .substring("2"),
  }),

  mounted() {
    this.tabs.register(this);
  },

  beforeDestroy() {
    this.tabs.unRegister(this);
  },

  computed: {
    tabs() {
      return this.$parent;
    },

    transition() {
      return this.tabs.transition;
    },

    isActived() {
      return this.activeModel === this.model && !this.disabled;
    },

    slideDirection() {
      if (!this.transition.slide) return "";

      if (this.transition.vertical) {
        return {
          left: "slide-top",
          right: "slide-bottom",
        }[this.transition.side];
      }

      return {
        left: "slide-left",
        right: "slide-right",
      }[this.transition.side];
    },

    ownNavItemIndex() {
      return this.tabs.navItems.findIndex(
        (navItem) => navItem.model === this.model
      );
    },
  },

  watch: {
    disabled(payload) {
      payload && this.tabs?.disableTabItem(this.ownNavItemIndex);
      this.tabs.navItems[this.ownNavItemIndex].disabled = payload;
    },
  },

  methods: {
    beforeEnter(el) {
      el.style.position = "relative";
    },

    leave(el) {
      el.style.position = "absolute";
    },
  },
};
</script>

<style scoped>
.tab-item {
  top: 0;
  left: 0;
  z-index: 1;
  width: 100%;
  height: inherit;
  background: white;
  transition: transform cubic-bezier(0.25, 0.8, 0.5, 1);
}

/* Horizontal */
.slide-left-enter,
.slide-right-leave-to {
  transform: translateX(-100%);
}

.slide-left-leave-to,
.slide-right-enter {
  transform: translateX(100%);
}

/* Vertical */
.slide-bottom-leave-to,
.slide-top-enter {
  transform: translateY(-100%);
}

.slide-top-leave-to,
.slide-bottom-enter {
  transform: translateY(100%);
}
</style>
