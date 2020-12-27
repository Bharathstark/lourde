import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import(/* webpackChunkName: "home" */ "../views/Home.vue")
  },
  {
    path: "/contact",
    name: "About",
    component: () => import(/* webpackChunkName: "home" */ "../views/About.vue")
  },
  {
    path: "/login",
    name: "Login",
    component: () =>
      import(/* webpackChunkName: "admin" */ "../views/Login.vue")
  },
  {
    path: "/admin",
    name: "Admin",
    meta: { requiresAuth: true },
    component: () =>
      import(/* webpackChunkName: "admin" */ "../views/Admin.vue")
  },
  {
    path: "/videos",
    name: "Videos",
    component: () =>
      import(/* webpackChunkName: "videos" */ "../views/Videos.vue")
  },
  {
    path: "/calendar",
    name: "Calendar",
    component: () =>
      import(/* webpackChunkName: "calendar" */ "../views/Calendar.vue")
  },
  {
    path: "/events",
    name: "Events",
    component: () =>
      import(/* webpackChunkName: "event" */ "../views/Events.vue")
  },
  {
    path: "/achievements",
    name: "Achievements",
    component: () =>
      import(/* webpackChunkName: "event" */ "../views/Events.vue")
  },
  {
    path: "/events/:id",
    name: "Event",
    component: () =>
      import(/* webpackChunkName: "event" */ "../views/Event.vue")
  },
  {
    path: "/achievements/:id",
    name: "Achievement",
    component: () =>
      import(/* webpackChunkName: "event" */ "../views/Event.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, _from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    console.log(Vue.$cookies.isKey("csrf_token"));
    if (!Vue.$cookies.isKey("csrf_token")) {
      next({
        path: "/login"
      });
    } else {
      next();
    }
  } else {
    next(); // make sure to always call next()!
  }
});

export default router;
