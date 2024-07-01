cmake_minimum_required(VERSION 3.16)
project(cattower)

find_package(Qt6 REQUIRED COMPONENTS Core Gui Qml Quick QuickControls2)

qt_standard_project_setup(REQUIRES 6.4)

qt_add_executable(cattower WIN32
    main.cpp
)

target_link_libraries(cattower PRIVATE
    Qt6::Core
    Qt6::Gui
    Qt6::Qml
    Qt6::Quick
    Qt6::QuickControls2
)

qt_add_qml_module(cattower
    URI nathanosman.cattower
    VERSION 1.0
    RESOURCE_PREFIX "/qt/qml"
    QML_FILES
        "Main.qml"
    RESOURCES
        "content/images/home.svg"
        "content/images/settings.svg"
        "content/images/wifi.svg"
)
