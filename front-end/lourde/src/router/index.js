import Home from "../views/Home.vue";
import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: { title: "Home",content:"This is the Home Page of Our Website" }
  },
  {
    path: "/index.html",
    component: Home,
    alias: '/',
    meta: { title: "Home",content:"This is the Home Page of Our Website" }
  },
  {
    path: "/contact",
    name: "About",
    meta: { title: "Contact Us",content:"Contact our Office Incase Of Queries" },
    component: () =>
      import(/* webpackChunkName: "contact" */ "../views/About.vue")
  },
  {
    path: "/login",
    name: "Login",
    meta: { title: "Login",content:"Login for Administrative Purposes" },
    component: () =>
      import(/* webpackChunkName: "login" */ "../views/Login.vue")
  },
  {
    path: "/admin",
    name: "Admin",
    meta: { requiresAuth: true,title: "Admin",content:"Admin DashBoard for Administrative Purposes" },
    component: () =>
      import(/* webpackChunkName: "admin" */ "../views/Admin.vue")
  },
  {
    path: "/videos",
    name: "Videos",
    meta: { title: "Videos",content:"Youtube Videos of School Events,Speeches etc" },
    component: () =>
      import(/* webpackChunkName: "videos" */ "../views/Videos.vue")
  },
  {
    path: "/calendar",
    name: "Calendar",
    meta: { title: "Calendar" ,content:"Administrative Calendar"},
    component: () =>
      import(/* webpackChunkName: "calendar" */ "../views/Calendar.vue")
  },
  {
    path: "/events",
    name: "Events",
    props: { module: "event" },
    meta: { title: "Events",content:"Events of the School" },
    component: () =>
      import(/* webpackChunkName: "events" */ "../views/Events.vue")
  },
  {
    path: "/achievements",
    name: "Achievements",
    props: { module: "achievement" },
    meta: { title: "Achievements",content:"Achievements of the School" },
    component: () =>
      import(/* webpackChunkName: "events" */ "../views/Events.vue")
  },
  {
    path: "/event",
    name: "Event",
    props: route => ({ eventID: route.query.id }),
    meta: { title: "Event",content:"An Event of the school with images" },
    component: () =>
      import(/* webpackChunkName: "event" */ "../views/Event.vue")
  },
  {
    path: "/achievement",
    name: "Achievement",
    props: route => ({ eventID: route.query.id }),
    meta: { title: "Achievement",content:"An Achievement of the school with images" },
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
  document.title=to.meta.title
  const tag = document.createElement('meta');
  tag.setAttribute('name','description');
  tag.setAttribute('content',to.meta.content);
  tag.setAttribute('data-vue-router-controlled','');
  document.head.appendChild(tag)
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!Vue.$cookies.isKey("csrf_token")) {
      next({
        path: "/login"
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
