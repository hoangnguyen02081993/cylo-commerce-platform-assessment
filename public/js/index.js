var filterNameValue = undefined
var filterPriceValue = undefined
var filterBranchValue = undefined
var sortFieldValue = undefined
var sortDirectionValue = 'asc'
let person = !!localStorage.getItem('person') ? localStorage.getItem('person') : (() => {
    const newPerson = (new Date()).getTime()
    localStorage.setItem('person', newPerson)
    return newPerson
})()

// event actions
const onChangePerson = () => {
    const newPerson = (new Date()).getTime()
    localStorage.setItem('person', newPerson)
    person = newPerson
}

const onFilterNameChanged = debounce(function (e) {
    filterNameValue = e.value
    filterProducts()
})

const onFilterPriceChanged = debounce(function (e) {
    filterPriceValue = e.value
    filterProducts()
})

const onFilterBranchChanged = debounce(function (e) {
    filterBranchValue = e.value
    filterProducts()
})

const onSortFieldChanged = debounce(function (e) {
    sortFieldValue = e.value
    filterProducts()
})

const onSortDirectionChanged = debounce(function (e) {
    sortDirectionValue = e.value
    filterProducts()
})
const gotoDetail = (id) => {
    window.location.href = `/detail.html?id=${id}`
}
// ------------------------------


// integrate functions
function buildQueries() {
    let queries = []
    if (filterNameValue) {
        queries.push(`name=${encodeURIComponent(filterNameValue)}`)
    }
    if (filterPriceValue) {
        if (filterPriceValue.includes('-')) {
            queries.push(`fromPrice=${filterPriceValue.split('-')[0]}`)
            queries.push(`toPrice=${filterPriceValue.split('-')[1]}`)
        }
        else if (filterPriceValue.endsWith('+')) {
            queries.push(`fromPrice=${filterPriceValue.replace('+', '')}`)
        }
    }
    if (filterBranchValue) {
        queries.push(`branch=${encodeURIComponent(filterBranchValue)}`)
    }

    if (sortFieldValue && sortDirectionValue) {
        queries.push(`sortField=${sortFieldValue}`)
        queries.push(`sortDirection=${sortDirectionValue}`)
    }

    return queries.join('&')
}
function filterProducts() {
    const queries = buildQueries()
    $.ajax({
        url: `/api/products/${queries ? '?' + queries : ''}`,
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
function renderProductHtml(products) {
    const elm = document.getElementById('product-list')
    if (elm) {
        if (!Array.isArray(products)) {
            elm.innerHTML = ''
        }
        else {
            elm.innerHTML = products.map(buildProductHtml).join('')
        }
    }
}
function buildProductHtml(product) {
    return `
    <li data-name="${product.name}" data-price="${product.price}" data-branch="${product.branc}">
        <img src="${product.imgUrl}" alt="${product.altName}">
        <h2>${product.name}</h2>
        <p>${product.description}</p>
        <p class="price">$${product.price}</p>
        <p class="branch">${product.branch}</p>
        <button onclick="gotoDetail(${product.id})">Go to detail</button>
    </li>
    `
}
// ------------------------------

document.addEventListener("DOMContentLoaded", () => {
    filterProducts()
});