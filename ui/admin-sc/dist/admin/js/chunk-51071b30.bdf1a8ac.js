(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-51071b30"],{"03a9":function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"give-badge-selection"},[r("div",{staticClass:"box dropdown"},[r("h2",{staticClass:"box-header"},[e._v("\n\t\t\tSelect badge to give\n\t\t")]),r("div",{staticClass:"box-content"},[r("ul",{staticClass:"selected-student"},e._l(e.selectStudent,(function(t,n){return r("li",{key:""+t+n,class:{error:e.errors[n]}},[r("a",{class:["dropdown",t.show?"active":""],on:{click:function(e){t.show=!t.show}}},[r("p",[e._v(e._s(t.studentId)+" "+e._s(t.fullName))]),r("icon-arrow-dropdown",{staticClass:"icon"})],1),r("transition",{attrs:{name:"slide-down"}},[t.show?r("div",{staticClass:"badge-form"},[r("div",{staticClass:"my-row"},e._l(e.options,(function(a,i){return r("b-col",{key:""+t.studentId+i,staticClass:"badge-wrapper",attrs:{lg:"3",cols:"6"}},[r("label",{class:["badge-checkbox",e.hasSelected(t.badges,a.id)?"is-select":""],attrs:{for:""+t.studentId+i}},[r("base-image",{attrs:{size:"90"}}),r("input",{directives:[{name:"model",rawName:"v-model",value:t.badges,expression:"item.badges"}],attrs:{id:""+t.studentId+i,type:"checkbox"},domProps:{value:a,checked:Array.isArray(t.badges)?e._i(t.badges,a)>-1:t.badges},on:{input:function(t){return e.removeError(n)},change:function(r){var n=t.badges,i=r.target,c=!!i.checked;if(Array.isArray(n)){var s=a,o=e._i(n,s);i.checked?o<0&&e.$set(t,"badges",n.concat([s])):o>-1&&e.$set(t,"badges",n.slice(0,o).concat(n.slice(o+1)))}else e.$set(t,"badges",c)}}}),r("p",{staticClass:"text"},[e._v(e._s(a.name))])],1)])})),1)]):e._e()])],1)})),0)])]),r("base-page-step",{attrs:{step:e.step},on:{next:e.submit,back:e.goBack}})],1)},a=[],i=(r("d93a"),r("ef92"),r("c1c3"),r("477f")),c=(r("5b54"),r("67c8"),r("a0c4"),r("e9de"),r("6b3c")),s=r("90b8"),o=r("94ea");function u(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function l(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?u(r,!0).forEach((function(t){Object(c["a"])(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):u(r).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}var d={components:{IconArrowDropdown:s["a"]},data:function(){return{selectStudent:[],errors:[],options:[{id:"002",name:"Team working"},{id:"003",name:"Communication"},{id:"004",name:"Leadership"},{id:"005",name:"Flexible"}]}},computed:l({},Object(o["b"])("giveBadge",["selectedStudents","steps"]),{step:function(){return this.$route.meta.step},hasError:function(){return this.errors.some((function(e){return e}))}}),beforeRouteEnter:function(e,t,r){r((function(e){e.steps.includes("main")||e.$router.replace({name:"give-badge"})}))},created:function(){this.selectStudent=this.selectedStudents},methods:{validateSubmit:function(){var e=this;this.errors=[],this.selectStudent.forEach((function(t){0===t.badges.length?e.errors.push(!0):e.errors.push(!1)}))},submit:function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.validateSubmit(),!this.hasError){e.next=4;break}return this.$bvToast.toast("Please select the badge to student",{title:"No badge error",variant:"danger",autoHideDelay:1500}),e.abrupt("return");case 4:return e.next=6,this.$store.dispatch("giveBadge/updateSelectedBadge",this.selectStudent);case 6:return e.next=8,this.$store.dispatch("giveBadge/addStep",this.step.step);case 8:this.$router.push({name:"give-badge-confirmation"});case 9:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),goBack:function(){this.$router.push({name:"give-badge"})},removeError:function(e){this.errors[e]=!1},hasSelected:function(e,t){return e.some((function(e){return e.id===t}))}}},f=d,p=(r("9003"),r("a6c2")),v=Object(p["a"])(f,n,a,!1,null,null,null);t["default"]=v.exports},"0e57":function(e,t,r){},"16ed":function(e,t,r){"use strict";var n=r("3855"),a=RegExp.prototype.exec;e.exports=function(e,t){var r=e.exec;if("function"===typeof r){var i=r.call(e,t);if("object"!==typeof i)throw new TypeError("RegExp exec method returned something other than an Object or null");return i}if("RegExp"!==n(e))throw new TypeError("RegExp#exec called on incompatible receiver");return a.call(e,t)}},3017:function(e,t,r){"use strict";var n=r("3c01");e.exports=function(){var e=n(this),t="";return e.global&&(t+="g"),e.ignoreCase&&(t+="i"),e.multiline&&(t+="m"),e.unicode&&(t+="u"),e.sticky&&(t+="y"),t}},"37af":function(e,t,r){"use strict";var n=r("9402");r("5fa5")({target:"RegExp",proto:!0,forced:n!==/./.exec},{exec:n})},"5a27":function(e,t,r){"use strict";r("37af");var n=r("b14f"),a=r("b639"),i=r("13b5"),c=r("1823"),s=r("f5dc"),o=r("9402"),u=s("species"),l=!i((function(){var e=/./;return e.exec=function(){var e=[];return e.groups={a:"7"},e},"7"!=="".replace(e,"$<a>")})),d=function(){var e=/(?:)/,t=e.exec;e.exec=function(){return t.apply(this,arguments)};var r="ab".split(e);return 2===r.length&&"a"===r[0]&&"b"===r[1]}();e.exports=function(e,t,r){var f=s(e),p=!i((function(){var t={};return t[f]=function(){return 7},7!=""[e](t)})),v=p?!i((function(){var t=!1,r=/a/;return r.exec=function(){return t=!0,null},"split"===e&&(r.constructor={},r.constructor[u]=function(){return r}),r[f](""),!t})):void 0;if(!p||!v||"replace"===e&&!l||"split"===e&&!d){var h=/./[f],g=r(c,f,""[e],(function(e,t,r,n,a){return t.exec===o?p&&!a?{done:!0,value:h.call(t,r,n)}:{done:!0,value:e.call(r,t,n)}:{done:!1}})),b=g[0],m=g[1];n(String.prototype,e,b),a(RegExp.prototype,f,2==t?function(e,t){return m.call(e,this,t)}:function(e){return m.call(e,this)})}}},"67c8":function(e,t,r){"use strict";var n=r("3c01"),a=r("5271"),i=r("4a9e"),c=r("9af8"),s=r("7558"),o=r("16ed"),u=Math.max,l=Math.min,d=Math.floor,f=/\$([$&`']|\d\d?|<[^>]*>)/g,p=/\$([$&`']|\d\d?)/g,v=function(e){return void 0===e?e:String(e)};r("5a27")("replace",2,(function(e,t,r,h){return[function(n,a){var i=e(this),c=void 0==n?void 0:n[t];return void 0!==c?c.call(n,i,a):r.call(String(i),n,a)},function(e,t){var a=h(r,e,this,t);if(a.done)return a.value;var d=n(e),f=String(this),p="function"===typeof t;p||(t=String(t));var b=d.global;if(b){var m=d.unicode;d.lastIndex=0}var x=[];while(1){var w=o(d,f);if(null===w)break;if(x.push(w),!b)break;var y=String(w[0]);""===y&&(d.lastIndex=s(f,i(d.lastIndex),m))}for(var S="",k=0,O=0;O<x.length;O++){w=x[O];for(var E=String(w[0]),$=u(l(c(w.index),f.length),0),j=[],_=1;_<w.length;_++)j.push(v(w[_]));var C=w.groups;if(p){var A=[E].concat(j,$,f);void 0!==C&&A.push(C);var R=String(t.apply(void 0,A))}else R=g(E,f,$,j,C,t);$>=k&&(S+=f.slice(k,$)+R,k=$+E.length)}return S+f.slice(k)}];function g(e,t,n,i,c,s){var o=n+e.length,u=i.length,l=p;return void 0!==c&&(c=a(c),l=f),r.call(s,l,(function(r,a){var s;switch(a.charAt(0)){case"$":return"$";case"&":return e;case"`":return t.slice(0,n);case"'":return t.slice(o);case"<":s=c[a.slice(1,-1)];break;default:var l=+a;if(0===l)return r;if(l>u){var f=d(l/10);return 0===f?r:f<=u?void 0===i[f-1]?a.charAt(1):i[f-1]+a.charAt(1):r}s=i[l-1]}return void 0===s?"":s}))}}))},7558:function(e,t,r){"use strict";var n=r("b910")(!0);e.exports=function(e,t,r){return t+(r?n(e,t).length:1)}},9003:function(e,t,r){"use strict";var n=r("0e57"),a=r.n(n);a.a},"90b8":function(e,t,r){"use strict";var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("svg",{attrs:{width:e.size,xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 320 512"}},[r("path",{attrs:{d:"M143 352.3L7 216.3c-9.4-9.4-9.4-24.6 0-33.9l22.6-22.6c9.4-9.4 24.6-9.4 33.9 0l96.4 96.4 96.4-96.4c9.4-9.4 24.6-9.4 33.9 0l22.6 22.6c9.4 9.4 9.4 24.6 0 33.9l-136 136c-9.2 9.4-24.4 9.4-33.8 0z"}})])},a=[],i={props:{size:{type:String,default:"12"}}},c=i,s=r("a6c2"),o=Object(s["a"])(c,n,a,!1,null,null,null);t["a"]=o.exports},9402:function(e,t,r){"use strict";var n=r("3017"),a=RegExp.prototype.exec,i=String.prototype.replace,c=a,s="lastIndex",o=function(){var e=/a/,t=/b*/g;return a.call(e,"a"),a.call(t,"a"),0!==e[s]||0!==t[s]}(),u=void 0!==/()??/.exec("")[1],l=o||u;l&&(c=function(e){var t,r,c,l,d=this;return u&&(r=new RegExp("^"+d.source+"$(?!\\s)",n.call(d))),o&&(t=d[s]),c=a.call(d,e),o&&c&&(d[s]=d.global?c.index+c[0].length:t),u&&c&&c.length>1&&i.call(c[0],r,(function(){for(l=1;l<arguments.length-2;l++)void 0===arguments[l]&&(c[l]=void 0)})),c}),e.exports=c},b910:function(e,t,r){var n=r("9af8"),a=r("1823");e.exports=function(e){return function(t,r){var i,c,s=String(a(t)),o=n(r),u=s.length;return o<0||o>=u?e?"":void 0:(i=s.charCodeAt(o),i<55296||i>56319||o+1===u||(c=s.charCodeAt(o+1))<56320||c>57343?e?s.charAt(o):i:e?s.slice(o,o+2):c-56320+(i-55296<<10)+65536)}}}}]);
//# sourceMappingURL=chunk-51071b30.bdf1a8ac.js.map