(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0bbbebb4"],{"4ad0":function(t,e,r){},"9c87":function(t,e,r){"use strict";var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("ol",{staticClass:"base-breadcrumb"},t._l(t.items,(function(e,a){return r("li",{key:""+e+a},[t._v("\n\t\t"+t._s(e)+"\n\t")])})),0)},s=[],n={props:{items:{type:Array,required:!0}}},c=n,u=(r("a63b"),r("a6c2")),i=Object(u["a"])(c,a,s,!1,null,null,null);e["a"]=i.exports},a63b:function(t,e,r){"use strict";var a=r("4ad0"),s=r.n(a);s.a},fb00:function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("section",{staticClass:"section"},[r("div",{staticClass:"page-header"},[r("h1",{staticClass:"title"},[t._v("\n\t\t\t"+t._s(t.currentPage)+"\n\t\t")]),r("base-breadcrumb",{staticClass:"breadcrumb",attrs:{items:t.breadcrumbList}})],1),r("router-view")],1)},s=[],n=r("9c87"),c={components:{BaseBreadcrumb:n["a"]},data:function(){return{breadcrumbList:[],isHideStep:!1}},computed:{currentPage:function(){return this.breadcrumbList[this.breadcrumbList.length-1]}},watch:{$route:function(){this.updateBreadcrumb()}},mounted:function(){this.updateBreadcrumb()},methods:{updateBreadcrumb:function(){this.breadcrumbList=this.$route.meta.breadcrumb}}},u=c,i=r("a6c2"),b=Object(i["a"])(u,a,s,!1,null,null,null);e["default"]=b.exports}}]);
//# sourceMappingURL=chunk-0bbbebb4.a6223cb6.js.map