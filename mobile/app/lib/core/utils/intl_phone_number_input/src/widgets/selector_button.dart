import 'package:celebut/core/utils/intl_phone_number_input/src/models/country_model.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/utils/selector_config.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/utils/test/test_helper.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/widgets/countries_search_list_widget.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/widgets/input_widget.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/widgets/item.dart';
import 'package:flutter/material.dart';

/// [SelectorButton]
class SelectorButton extends StatelessWidget {
  final List<Country> countries;
  final Country? country;
  final SelectorConfig selectorConfig;
  final TextStyle? selectorTextStyle;
  final InputDecoration? searchBoxDecoration;
  final bool autoFocusSearchField;
  final String? locale;
  final bool isEnabled;
  final bool isScrollControlled;
  final VoidCallback onPressed;

  final ValueChanged<Country?> onCountryChanged;

  const SelectorButton({
    Key? key,
    required this.countries,
    required this.country,
    required this.selectorConfig,
    required this.selectorTextStyle,
    required this.searchBoxDecoration,
    required this.autoFocusSearchField,
    required this.locale,
    required this.onPressed,
    required this.onCountryChanged,
    required this.isEnabled,
    required this.isScrollControlled,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
          border: Border.all(), borderRadius: BorderRadius.circular(8)),
      child: MaterialButton(
        key: const Key(TestHelper.dropdownButtonKeyValue),
        padding: EdgeInsets.zero,
        minWidth: 0,
        onPressed: countries.isNotEmpty && countries.length > 1 && isEnabled
            ? () async {
                Country? selected;
                if (selectorConfig.selectorType ==
                    PhoneInputSelectorType.bottomSheet) {
                  selected =
                      await showCountrySelectorBottomSheet(context, countries);
                }
                if (selected != null) {
                  onCountryChanged(selected);
                }
              }
            : null,
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 6),
          child: Item(
            country: country,
            showFlag: selectorConfig.showFlags,
            useEmoji: selectorConfig.useEmoji,
            leadingPadding: selectorConfig.leadingPadding,
            trailingSpace: selectorConfig.trailingSpace,
            textStyle: selectorTextStyle,
          ),
        ),
      ),
    );
  }

  /// shows a Dialog with list [countries] if the [PhoneInputSelectorType.BOTTOM_SHEET] is selected
  Future<Country?> showCountrySelectorBottomSheet(
      BuildContext inheritedContext, List<Country> countries) {
    return showModalBottomSheet(
      context: inheritedContext,
      clipBehavior: Clip.hardEdge,
      isScrollControlled: isScrollControlled,
      backgroundColor: Colors.transparent,
      shape: const RoundedRectangleBorder(
          borderRadius: BorderRadius.only(
              topLeft: Radius.circular(12), topRight: Radius.circular(12))),
      useSafeArea: selectorConfig.useBottomSheetSafeArea,
      builder: (BuildContext context) {
        return Stack(children: [
          GestureDetector(
            onTap: () => Navigator.pop(context),
          ),
          Padding(
            padding: EdgeInsets.only(
                bottom: MediaQuery.of(context).viewInsets.bottom),
            child: DraggableScrollableSheet(
              builder: (BuildContext context, ScrollController controller) {
                return Directionality(
                  textDirection: Directionality.of(inheritedContext),
                  child: Container(
                    padding:
                        const EdgeInsets.only(top: 15, left: 10, right: 10),
                    decoration: const ShapeDecoration(
                      color: Colors.white,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.only(
                          topLeft: Radius.circular(12),
                          topRight: Radius.circular(12),
                        ),
                      ),
                    ),
                    child: CountrySearchListWidget(
                      countries,
                      locale,
                      searchBoxDecoration: searchBoxDecoration,
                      scrollController: controller,
                      showFlags: selectorConfig.showFlags,
                      useEmoji: selectorConfig.useEmoji,
                      autoFocus: autoFocusSearchField,
                    ),
                  ),
                );
              },
            ),
          ),
        ]);
      },
    );
  }
}
