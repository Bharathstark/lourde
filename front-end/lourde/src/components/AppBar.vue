<template>
  <v-menu top v-bind="$attrs" v-on="$listeners">
    <v-list>
      <v-list-item v-for="(AppBarItem, index) in AppBarItems" :key="index">
        <v-btn text :href="AppBarItem.href">{{ AppBarItem.display }}</v-btn>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script>
export default {
  name: "AppBar",
  data: () => ({
    drawer: false,
    AppBarItems: [
      {
        display: "Home",
        href: "/"
      },
      {
        display: "Achievements",
        href: "/achievements"
      },
      {
        display: "Events",
        href: "/events"
      },
      {
        display: "Videos",
        href: "/videos"
      },
      {
        display: "Calendar",
        href: "/calendar"
      },
      {
        display: "Contact Us",
        href: "/contact"
      }
    ]
  })
};
</script>

<style></style>
