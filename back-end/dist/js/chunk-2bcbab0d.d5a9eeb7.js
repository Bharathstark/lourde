(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2bcbab0d"],{1800:function(t,e,i){"use strict";i("4de4");var s=i("2b0e");e["a"]=s["a"].extend({name:"v-list-item-action",functional:!0,render:function(t,e){var i=e.data,s=e.children,a=void 0===s?[]:s;i.staticClass=i.staticClass?"v-list-item__action ".concat(i.staticClass):"v-list-item__action";var n=a.filter((function(t){return!1===t.isComment&&" "!==t.text}));return n.length>1&&(i.staticClass+=" v-list-item__action--stack"),t("div",i,a)}})},"1bfb":function(t,e,i){},"34c3":function(t,e,i){"use strict";i("498a");var s=i("2b0e");e["a"]=s["a"].extend({name:"v-list-item-icon",functional:!0,render:function(t,e){var i=e.data,s=e.children;return i.staticClass="v-list-item__icon ".concat(i.staticClass||"").trim(),t("div",i,s)}})},"5a70":function(t,e,i){"use strict";i.r(e);var s=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("section",[i("div",{staticClass:"py-6"}),i("h2",{staticClass:"text-md-h3 text-sm-h3  mb-3 text-center "},[t._v("Academics")]),i("div",{staticClass:"py-6"}),i("v-card",{attrs:{flat:""}},[i("v-tabs",{attrs:{color:"deep-purple accent-4",centered:""}},[i("v-tab",[t._v("12 Standard Results")]),i("v-tab",[t._v("10 Standard Results")]),t._l(t.studentData,(function(e,s){return i("v-tab-item",{key:s},[i("v-container",[i("v-row",t._l(e,(function(e){return i("v-col",{key:e.id,staticClass:"d-flex flex-column",attrs:{cols:"12",xs:"6",sm:"4",md:"3"}},[i("v-card",{staticClass:" flex d-flex flex-column"},[i("div",{staticClass:" mt-8 text-center"},[i("v-avatar",{attrs:{size:"130"}},[i("v-img",{attrs:{large:"",src:e.thumbnail}})],1)],1),i("v-list-item",[i("v-list-item-content",[i("v-list-item-title",{staticClass:"text-deep-purple accent-4 ma-2 text-center"},[t._v(t._s(e.name))]),i("v-list-item-subtitle",{staticClass:" text-center"},[t._v(t._s(e.description))])],1)],1),t._l(e.marks,(function(e){return i("v-list-item",{key:e.id,staticClass:"flex ml-14"},[i("v-list-item-content",[t._v(t._s(e.field))]),i("v-list-item-content",{staticClass:"ml-2"},[t._v(" "+t._s(e.value)+" ")])],1)}))],2)],1)})),1)],1)],1)}))],2)],1)],1)},a=[],n=(i("96cf"),i("1da1")),r={name:"Academics",data:function(){return{studentURL:"/api/v1/students",studentData:[]}},created:function(){var t=this;this.fetchData(this.studentURL+"?module=12").then((function(e){return t.studentData.push(e)})),this.fetchData(this.studentURL+"?module=10").then((function(e){return t.studentData.push(e)}))},methods:{fetchData:function(t){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function i(){return regeneratorRuntime.wrap((function(i){while(1)switch(i.prev=i.next){case 0:return i.abrupt("return",e.axios.get(t).then((function(t){return t.data["data"]})).catch((function(t){console.log(t)})));case 1:case"end":return i.stop()}}),i)})))()}}},o=r,l=i("2877"),c=i("6544"),u=i.n(c),d=i("8212"),h=i("b0af"),v=i("62ad"),p=i("a523"),f=i("adda"),m=i("da13"),b=i("5d23"),g=i("0fd9"),C=i("71a3"),x=i("c671"),_=i("fe57"),w=Object(l["a"])(o,s,a,!1,null,null,null);e["default"]=w.exports;u()(w,{VAvatar:d["a"],VCard:h["a"],VCol:v["a"],VContainer:p["a"],VImg:f["a"],VListItem:m["a"],VListItemContent:b["a"],VListItemSubtitle:b["b"],VListItemTitle:b["c"],VRow:g["a"],VTab:C["a"],VTabItem:x["a"],VTabs:_["a"]})},"5d23":function(t,e,i){"use strict";i.d(e,"a",(function(){return S})),i.d(e,"c",(function(){return j})),i.d(e,"b",(function(){return k}));var s=i("80d2"),a=i("8860"),n=(i("ac1f"),i("466d"),i("5530")),r=i("ade3"),o=(i("db42"),i("9d26")),l=i("da13"),c=i("34c3"),u=i("7e2b"),d=i("9d65"),h=i("a9ad"),v=i("f2e7"),p=i("3206"),f=i("5607"),m=i("0789"),b=i("58df"),g=Object(b["a"])(u["a"],d["a"],h["a"],Object(p["a"])("list"),v["a"]),C=g.extend().extend({name:"v-list-group",directives:{ripple:f["a"]},props:{activeClass:{type:String,default:""},appendIcon:{type:String,default:"$expand"},color:{type:String,default:"primary"},disabled:Boolean,group:String,noAction:Boolean,prependIcon:String,ripple:{type:[Boolean,Object],default:!0},subGroup:Boolean},computed:{classes:function(){return{"v-list-group--active":this.isActive,"v-list-group--disabled":this.disabled,"v-list-group--no-action":this.noAction,"v-list-group--sub-group":this.subGroup}}},watch:{isActive:function(t){!this.subGroup&&t&&this.list&&this.list.listClick(this._uid)},$route:"onRouteChange"},created:function(){this.list&&this.list.register(this),this.group&&this.$route&&null==this.value&&(this.isActive=this.matchRoute(this.$route.path))},beforeDestroy:function(){this.list&&this.list.unregister(this)},methods:{click:function(t){var e=this;this.disabled||(this.isBooted=!0,this.$emit("click",t),this.$nextTick((function(){return e.isActive=!e.isActive})))},genIcon:function(t){return this.$createElement(o["a"],t)},genAppendIcon:function(){var t=!this.subGroup&&this.appendIcon;return t||this.$slots.appendIcon?this.$createElement(c["a"],{staticClass:"v-list-group__header__append-icon"},[this.$slots.appendIcon||this.genIcon(t)]):null},genHeader:function(){return this.$createElement(l["a"],{staticClass:"v-list-group__header",attrs:{"aria-expanded":String(this.isActive),role:"button"},class:Object(r["a"])({},this.activeClass,this.isActive),props:{inputValue:this.isActive},directives:[{name:"ripple",value:this.ripple}],on:Object(n["a"])(Object(n["a"])({},this.listeners$),{},{click:this.click})},[this.genPrependIcon(),this.$slots.activator,this.genAppendIcon()])},genItems:function(){var t=this;return this.showLazyContent((function(){return[t.$createElement("div",{staticClass:"v-list-group__items",directives:[{name:"show",value:t.isActive}]},Object(s["o"])(t))]}))},genPrependIcon:function(){var t=this.subGroup&&null==this.prependIcon?"$subgroup":this.prependIcon;return t||this.$slots.prependIcon?this.$createElement(c["a"],{staticClass:"v-list-group__header__prepend-icon"},[this.$slots.prependIcon||this.genIcon(t)]):null},onRouteChange:function(t){if(this.group){var e=this.matchRoute(t.path);e&&this.isActive!==e&&this.list&&this.list.listClick(this._uid),this.isActive=e}},toggle:function(t){var e=this,i=this._uid===t;i&&(this.isBooted=!0),this.$nextTick((function(){return e.isActive=i}))},matchRoute:function(t){return null!==t.match(this.group)}},render:function(t){return t("div",this.setTextColor(this.isActive&&this.color,{staticClass:"v-list-group",class:this.classes}),[this.genHeader(),t(m["a"],this.genItems())])}}),x=(i("899c"),i("604c")),_=Object(b["a"])(x["a"],h["a"]).extend({name:"v-list-item-group",provide:function(){return{isInGroup:!0,listItemGroup:this}},computed:{classes:function(){return Object(n["a"])(Object(n["a"])({},x["a"].options.computed.classes.call(this)),{},{"v-list-item-group":!0})}},methods:{genData:function(){return this.setTextColor(this.color,Object(n["a"])(Object(n["a"])({},x["a"].options.methods.genData.call(this)),{},{attrs:{role:"listbox"}}))}}}),w=i("1800"),O=i("8270"),$=Object(s["h"])("v-list-item__action-text","span"),S=Object(s["h"])("v-list-item__content","div"),j=Object(s["h"])("v-list-item__title","div"),k=Object(s["h"])("v-list-item__subtitle","div");a["a"],l["a"],w["a"],O["a"],c["a"]},"71a3":function(t,e,i){"use strict";i("c975"),i("ac1f"),i("5319");var s=i("5530"),a=i("4e82"),n=i("1c87"),r=i("7560"),o=i("80d2"),l=i("58df"),c=Object(l["a"])(n["a"],Object(a["a"])("tabsBar"),r["a"]);e["a"]=c.extend().extend().extend({name:"v-tab",props:{ripple:{type:[Boolean,Object],default:!0}},data:function(){return{proxyClass:"v-tab--active"}},computed:{classes:function(){return Object(s["a"])(Object(s["a"])({"v-tab":!0},n["a"].options.computed.classes.call(this)),{},{"v-tab--disabled":this.disabled},this.groupClasses)},value:function(){var t=this.to||this.href||"";if(this.$router&&this.to===Object(this.to)){var e=this.$router.resolve(this.to,this.$route,this.append);t=e.href}return t.replace("#","")}},mounted:function(){this.onRouteChange()},methods:{click:function(t){this.href&&this.href.indexOf("#")>-1&&t.preventDefault(),t.detail&&this.$el.blur(),this.$emit("click",t),this.to||this.toggle()}},render:function(t){var e=this,i=this.generateRouteLink(),a=i.tag,n=i.data;return n.attrs=Object(s["a"])(Object(s["a"])({},n.attrs),{},{"aria-selected":String(this.isActive),role:"tab",tabindex:0}),n.on=Object(s["a"])(Object(s["a"])({},n.on),{},{keydown:function(t){t.keyCode===o["t"].enter&&e.click(t),e.$emit("keydown",t)}}),t(a,n,this.$slots.default)}})},8270:function(t,e,i){"use strict";i("a9e3");var s=i("5530"),a=i("713a");e["a"]=a["a"].extend({name:"v-list-item-avatar",props:{horizontal:Boolean,size:{type:[Number,String],default:40}},computed:{classes:function(){return Object(s["a"])(Object(s["a"])({"v-list-item__avatar--horizontal":this.horizontal},a["a"].options.computed.classes.call(this)),{},{"v-avatar--tile":this.tile||this.horizontal})}},render:function(t){var e=a["a"].options.render.call(this,t);return e.data=e.data||{},e.data.staticClass+=" v-list-item__avatar",e}})},"899c":function(t,e,i){},c671:function(t,e,i){"use strict";var s=i("1e6c");e["a"]=s["a"].extend({name:"v-tab-item",props:{id:String},methods:{genWindowItem:function(){var t=s["a"].options.methods.genWindowItem.call(this);return t.data.domProps=t.data.domProps||{},t.data.domProps.id=this.id||this.value,t}}})},db42:function(t,e,i){},fe57:function(t,e,i){"use strict";i("b0c0"),i("a9e3");var s=i("5530"),a=(i("1bfb"),i("b85c")),n=i("7efd"),r=i("7560"),o=i("d10f"),l=i("58df"),c=Object(l["a"])(n["a"],o["a"],r["a"]).extend({name:"v-tabs-bar",provide:function(){return{tabsBar:this}},computed:{classes:function(){return Object(s["a"])(Object(s["a"])({},n["a"].options.computed.classes.call(this)),{},{"v-tabs-bar":!0,"v-tabs-bar--is-mobile":this.isMobile,"v-tabs-bar--show-arrows":this.showArrows},this.themeClasses)}},watch:{items:"callSlider",internalValue:"callSlider",$route:"onRouteChange"},methods:{callSlider:function(){this.isBooted&&this.$emit("call:slider")},genContent:function(){var t=n["a"].options.methods.genContent.call(this);return t.data=t.data||{},t.data.staticClass+=" v-tabs-bar__content",t},onRouteChange:function(t,e){if(!this.mandatory){var i,s=this.items,n=t.path,r=e.path,o=!1,l=!1,c=Object(a["a"])(s);try{for(c.s();!(i=c.n()).done;){var u=i.value;if(u.to===n?o=!0:u.to===r&&(l=!0),o&&l)break}}catch(d){c.e(d)}finally{c.f()}!o&&l&&(this.internalValue=void 0)}}},render:function(t){var e=n["a"].options.render.call(this,t);return e.data.attrs={role:"tablist"},e}}),u=i("f665"),d=i("604c"),h=u["a"].extend({name:"v-tabs-items",props:{mandatory:{type:Boolean,default:!1}},computed:{classes:function(){return Object(s["a"])(Object(s["a"])({},u["a"].options.computed.classes.call(this)),{},{"v-tabs-items":!0})},isDark:function(){return this.rootIsDark}},methods:{getValue:function(t,e){return t.id||d["a"].options.methods.getValue.call(this,t,e)}}}),v=i("a9ad"),p=Object(l["a"])(v["a"]).extend({name:"v-tabs-slider",render:function(t){return t("div",this.setBackgroundColor(this.color,{staticClass:"v-tabs-slider"}))}}),f=i("a452"),m=i("dc22"),b=i("80d2"),g=Object(l["a"])(v["a"],f["a"],r["a"]);e["a"]=g.extend().extend({name:"v-tabs",directives:{Resize:m["a"]},props:{activeClass:{type:String,default:""},alignWithTitle:Boolean,backgroundColor:String,centerActive:Boolean,centered:Boolean,fixedTabs:Boolean,grow:Boolean,height:{type:[Number,String],default:void 0},hideSlider:Boolean,iconsAndText:Boolean,mobileBreakpoint:[String,Number],nextIcon:{type:String,default:"$next"},optional:Boolean,prevIcon:{type:String,default:"$prev"},right:Boolean,showArrows:[Boolean,String],sliderColor:String,sliderSize:{type:[Number,String],default:2},vertical:Boolean},data:function(){return{resizeTimeout:0,slider:{height:null,left:null,right:null,top:null,width:null},transitionTime:300}},computed:{classes:function(){return Object(s["a"])({"v-tabs--align-with-title":this.alignWithTitle,"v-tabs--centered":this.centered,"v-tabs--fixed-tabs":this.fixedTabs,"v-tabs--grow":this.grow,"v-tabs--icons-and-text":this.iconsAndText,"v-tabs--right":this.right,"v-tabs--vertical":this.vertical},this.themeClasses)},isReversed:function(){return this.$vuetify.rtl&&this.vertical},sliderStyles:function(){return{height:Object(b["f"])(this.slider.height),left:this.isReversed?void 0:Object(b["f"])(this.slider.left),right:this.isReversed?Object(b["f"])(this.slider.right):void 0,top:this.vertical?Object(b["f"])(this.slider.top):void 0,transition:null!=this.slider.left?null:"none",width:Object(b["f"])(this.slider.width)}},computedColor:function(){return this.color?this.color:this.isDark&&!this.appIsDark?"white":"primary"}},watch:{alignWithTitle:"callSlider",centered:"callSlider",centerActive:"callSlider",fixedTabs:"callSlider",grow:"callSlider",right:"callSlider",showArrows:"callSlider",vertical:"callSlider","$vuetify.application.left":"onResize","$vuetify.application.right":"onResize","$vuetify.rtl":"onResize"},mounted:function(){var t=this;this.$nextTick((function(){window.setTimeout(t.callSlider,30)}))},methods:{callSlider:function(){var t=this;return!this.hideSlider&&this.$refs.items&&this.$refs.items.selectedItems.length?(this.$nextTick((function(){var e=t.$refs.items.selectedItems[0];if(!e||!e.$el)return t.slider.width=0,void(t.slider.left=0);var i=e.$el;t.slider={height:t.vertical?i.scrollHeight:Number(t.sliderSize),left:t.vertical?0:i.offsetLeft,right:t.vertical?0:i.offsetLeft+i.offsetWidth,top:i.offsetTop,width:t.vertical?Number(t.sliderSize):i.scrollWidth}})),!0):(this.slider.width=0,!1)},genBar:function(t,e){var i=this,s={style:{height:Object(b["f"])(this.height)},props:{activeClass:this.activeClass,centerActive:this.centerActive,dark:this.dark,light:this.light,mandatory:!this.optional,mobileBreakpoint:this.mobileBreakpoint,nextIcon:this.nextIcon,prevIcon:this.prevIcon,showArrows:this.showArrows,value:this.internalValue},on:{"call:slider":this.callSlider,change:function(t){i.internalValue=t}},ref:"items"};return this.setTextColor(this.computedColor,s),this.setBackgroundColor(this.backgroundColor,s),this.$createElement(c,s,[this.genSlider(e),t])},genItems:function(t,e){var i=this;return t||(e.length?this.$createElement(h,{props:{value:this.internalValue},on:{change:function(t){i.internalValue=t}}},e):null)},genSlider:function(t){return this.hideSlider?null:(t||(t=this.$createElement(p,{props:{color:this.sliderColor}})),this.$createElement("div",{staticClass:"v-tabs-slider-wrapper",style:this.sliderStyles},[t]))},onResize:function(){this._isDestroyed||(clearTimeout(this.resizeTimeout),this.resizeTimeout=window.setTimeout(this.callSlider,0))},parseNodes:function(){for(var t=null,e=null,i=[],s=[],a=this.$slots.default||[],n=a.length,r=0;r<n;r++){var o=a[r];if(o.componentOptions)switch(o.componentOptions.Ctor.options.name){case"v-tabs-slider":e=o;break;case"v-tabs-items":t=o;break;case"v-tab-item":i.push(o);break;default:s.push(o)}else s.push(o)}return{tab:s,slider:e,items:t,item:i}}},render:function(t){var e=this.parseNodes(),i=e.tab,s=e.slider,a=e.items,n=e.item;return t("div",{staticClass:"v-tabs",class:this.classes,directives:[{name:"resize",modifiers:{quiet:!0},value:this.onResize}]},[this.genBar(i,s),this.genItems(a,n)])}})}}]);
//# sourceMappingURL=chunk-2bcbab0d.d5a9eeb7.js.map