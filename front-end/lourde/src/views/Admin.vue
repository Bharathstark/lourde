<template>
  <section>
    <component
      class="ml-15"
      :key="bindinfo.module"
      :is="dash"
      :bindinfo="bindinfo"
      @showEvents="showEvents"
    />

    <v-navigation-drawer
      v-model="drawer"
      expand-on-hover
      permanent
      fixed
      style="z-index:1;"
    >
      <v-list-item class="px-2">
        <v-list-item-avatar>
          <v-icon>mdi-chevron-left</v-icon>
        </v-list-item-avatar>

        <v-list-item-title>{{
          this.$cookies.get("username")
        }}</v-list-item-title>

        <v-btn icon @click.stop="mini = !mini">
          <v-icon>mdi-chevron-left</v-icon>
        </v-btn>
      </v-list-item>

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          @click="changeModule(item)"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <alertDialog ref="showed"></alertDialog>
  </section>
</template>

<script>
import alertDialog from "@/components/SnackBarDialog.vue";
export default {
  name: "Admin",
  components: {
    DashBoard: () => import("@/components/DashBoard.vue"),
    AdminEvent: () => import("@/components/AdminEvent.vue"),
    AdminEvents: () => import("@/components/AdminEntities.vue"),
    AdminCalendar: () => import("@/components/AdminCalendar.vue"),
    AdminVideos: () => import("@/components/AdminVideos.vue"),
    AdminStudents: () => import("@/components/AdminStudents.vue"),
    AdminPrincipal: () => import("@/components/AdminPrincipal.vue"),
    alertDialog
  },
  data: () => ({
    dash: "DashBoard",
    drawer: true,
    bindinfo: { module: "DashBoard" },
    items: [
      {
        title: "DashBoard",
        icon: "mdi-home-city",
        component: "DashBoard",
        bindinfo: { module: "DashBoard" }
      },
      {
        title: "Events",
        icon: "mdi-calendar-range-outline",
        component: "AdminEvent",
        bindinfo: { module: "Events" }
      },
      {
        title: "Achievements",
        icon: "mdi-trophy",
        component: "AdminEvent",
        bindinfo: { module: "Achievements" }
      },
      {
        title: "Calendar",
        icon: "mdi-calendar",
        component: "AdminCalendar",
        bindinfo: { module: "Calendar" }
      },
      {
        title: "Videos",
        icon: "mdi-youtube",
        component: "AdminVideos",
        bindinfo: { module: "AdminVideos" }
      },
      {
        title: "Students",
        icon: "mdi-account-group-outline",
        component: "AdminStudents",
        bindinfo: { module: "AdminStudents" }
      },
      {
        title: "Principal",
        icon: "mdi-account",
        component: "AdminPrincipal",
        bindinfo: { module: "AdminPrincipal" }
      }
    ],
    mini: true
  }),
  methods: {
    changeModule(value) {
      this.mini = !this.mini;
      this.dash = value.component;
      this.bindinfo = value.bindinfo;
    },
    showEvents(value) {
      this.dash = "AdminEvents";
      this.bindinfo = value;
    }
  },
  mounted() {
    var vm = this.$refs.showed.open;
    this.axios.interceptors.response.use(
      response => {
        if (response.data.response_status.response_code == 401) {
          vm("Session Expired, kindly Relogin - Click here ", "red", -1, true);
        }
        return response;
      },
      function(error) {
        // Do something with response error
        return Promise.reject(error);
      }
    );
  }
};
</script>

<style></style>
