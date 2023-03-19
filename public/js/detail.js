const person = !!localStorage.getItem('person') ? localStorage.getItem('person') : (() => {
    const newPerson = (new Date()).getTime()
    localStorage.setItem('person', newPerson)
    return newPerson
})()

function goBack() {
    window.location.href = '/'
}

function getDetail(id) {
    $.ajax({
        url: `/api/products/${id}`,
        headers: {
            'X-Person': person,
            'Content-Type': 'application/json'
        },
        method: 'GET',
        dataType: 'json',
        success: function (data, success) {
            if (success === "success") {
                if (data.result === 0) {
                    renderProductHtml(data.data)
                    Toastify({
                        text: 'Get successfully',
                    }).showToast();
                }
                else {
                    Toastify({
                        text: data.errorMesssage,
                        style: {
                            background: "linear-gradient(to right, #ff0100, #ff1200)",
                        },
                    }).showToast();

                }
            }
            else {
                Toastify({
                    text: data.errorMesssage ?? 'You need to login first',
                    style: {
                        background: "linear-gradient(to right, #ff0100, #ff1200)",
                    },
                }).showToast();
            }
        }
    });
}

// ------------------------------

// render functions
function renderProductHtml(product) {
    const elm = document.getElementById('product-details')
    if (elm) {
        elm.innerHTML = buildProductHtml(product)
    }
}
function buildProductHtml(product) {
    return `
        <h1>${product.name}</h1>
        <img src="${product.imgUrl}" alt="${product.name}">
        <p><strong>Price:</strong> $${product.price}</p>
        <p><strong>Branch:</strong> ${product.branch}</p>
        <p><strong>Description:</strong> ${product.description}</p>
        <button onclick="goBack()">Back</button>
    `
}
// ------------------------------

document.addEventListener("DOMContentLoaded", () => {
    const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
    });
    let id = params.id;
    if (!id) {
        goBack()
        return
    }

    getDetail(id)
});