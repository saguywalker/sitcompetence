const appHey = document.body;
const node = document.createElement("div");
node.setAttribute("class", "app-loader");
node.setAttribute("id", "spin-loader");

const loadme = {};

loadme.start = () => appHey.appendChild(node);

loadme.stop = () => appHey.removeChild(node);

export default loadme;