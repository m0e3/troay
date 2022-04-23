$(document).ready(() => {
    setSidebarHeight();
    $(window).on('resize', () => setSidebarHeight());
    $('.toggler').on('click', () => {
        $('.menu-container').toggleClass('active');
    });
    $('.nav-toggler').on('click', () => {
        $('.navbar-toggler').toggleClass('is-active');
        $('.navbar-menu').toggleClass('is-active');
    });
});
function setSidebarHeight () {
    const navbarHeight = $('.navbar').outerHeight();
    $('.menu-warpper')
        .outerHeight(document.documentElement.clientHeight - navbarHeight)
    .niceScroll({ emulatetouch:true });
}

