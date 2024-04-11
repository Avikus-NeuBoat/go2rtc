project(
  go2rtc
  LANGUAGES Go)

include(golang)

add_go_executable(go2rtc)

install(PROGRAMS ${CMAKE_CURRENT_BINARY_DIR}/go2rtc
        DESTINATION bin)
install(FILES ${CMAKE_CURRENT_SOURCE_DIR}/go2rtc.yaml
        DESTINATION ${CMAKE_CFG_PATH})