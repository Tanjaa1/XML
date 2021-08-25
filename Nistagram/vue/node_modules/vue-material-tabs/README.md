# vue-material-tabs

Vue 2 material tabs component ❤

### [Demo](https://vue-material-tabs.vercel.app/)

## ⚓ Features

- Nav with dynamic pagination.
- Touchlable pagination.
- Vertical support.
- Very flexible and customizable
- Multiple style themes.
- Zero dependencies

## 🚚 Install

```c
 yarn add vue-material-tabs  // npm install --save vue-material-tabs
```

## 🚀 Usage

### Global

```javascript
import Vue from "vue";
import Tabs from "vue-material-tabs";

Vue.use(Tabs);
```

### Local

```javascript
import { tabs, tab } from "vue-material-tabs";

export default {
  components: {
    tabs,
    tabItem,
  },
};
```

## 📌 Examples

```javascript
<Tabs>
  <TabItem name="Foo">
    <div class="first-tab">First tab</div>
  </TabItem>
  <TabItem name="Bar">
    <div class="second-tab">Second tab</div>
  </TabItem>
</Tabs>
```

#### You can customize the name in the nav by slot #name (This removes the use of the prop name)

```javascript
<Tabs>
  <TabItem>
    <template #name>
        My custom title 🍉
    </template>

    <div class="first-tab">
    First tab
    </div>
  </TabItem>
</Tabs>
```

### 👑 Themes

There are some themes available to customize your tab, you can apply them through the "theme" prop.

##### The default theme is, guess what, default, curious isn't it?!

- default
- purple
- red
- pink
- cyan
- green

#### Example

```python
<Tabs theme="cyan" >
  ....
</Tabs>
```

You can create your own personalized themes, passing an object with the properties through the "theme" prop.

#### Example

```python
<Tabs :theme="theme" >
  ....
</Tabs>
....
<script>
  export default {
    data: () => ({
      theme: {
        nav: "#4A148C",
        navItem: "#BDBDBD",
        navActiveItem: "#fff",
        slider: "#CE93D8",
        arrow: "#f3f3f3",
      },
    }),
  };
</script>
```

## Tabs

### Props

| Name          | type             | Default   | description                                       |
| ------------- | ---------------- | --------- | ------------------------------------------------- |
| value         | String           | undefined | The designated model value for the component.     |
| theme         | String - Object  | 'default' | Apply a custom theme.                             |
| vertical      | Boolean          | false     | Uses a vertical transition when changing windows. |
| ripple        | Boolean          | true      | Enalbe/disable ripple buttons effects.            |
| slideDuration | String - Number  | 200       | Set time in ms slide duration.                    |
| slideVertical | Boolean          | false     | Enable vertical slide animation.                  |
| slide         | Boolean - Object | false     | Enable/disable slide or set object with props.    |
| navAuto       | Boolean          | false     | Set nav auto items.                               |
| navSlider     | Boolean          | true      | Enable/disable slider under nav item.             |
| noTouch       | Boolean          | false     | Enable/disable slider tabs by touch.              |

### Slots

| Name | Description                   | Props                               |
| ---- | ----------------------------- | ----------------------------------- |
| nav  | Slot to replace the nav menu. | { navItems: Array, active: String } |

### Events

| Name  | Description                  |
| ----- | ---------------------------- |
| input | Emitted when tab is changed. |

## TabItem

### Props

| Name     | type    | Default     | description                                           |
| -------- | ------- | ----------- | ----------------------------------------------------- |
| name     | String  | 'Tab Item ' | Sets the tab value in the nav menu.                   |
| disabled | Boolean | false       | Removes the ability to click or target the component. |

### Slots

| Name | Description                |
| ---- | -------------------------- |
| name | Slot for cutom name in nav |

## 🔖 License

[MIT](https://github.com/jairoblatt/vue-material-tabs/blob/main/LICENSE)
