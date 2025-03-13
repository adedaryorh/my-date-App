import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class CustomRadioListTile extends StatelessWidget {
  const CustomRadioListTile({
    required this.option,
    required this.selectedOption,
    required this.onSelected,
    super.key,
  });

  final String option;
  final String? selectedOption;
  final ValueChanged<String> onSelected;

  @override
  Widget build(BuildContext context) {
    final isSelected = selectedOption == option;
    return GestureDetector(
      onTap: () => onSelected(option),
      child: Container(
        height: 38,
        width: double.maxFinite,
        margin: const EdgeInsets.only(bottom: 10),
        padding: const EdgeInsets.symmetric(horizontal: 10),
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(Radius.circular(8)),
          color: Color(0xffF4F4F4),
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(
              children: [
                const Space(5),
                Text(
                  option,
                  style: context.textTheme.bodyMedium,
                ),
              ],
            ),
            Container(
              height: 15.7,
              width: 13,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                border: Border.all(
                  width: isSelected ? 5 : 1,
                  color: isSelected
                      ? context.colorScheme.primary
                      : const Color(0xff96A7AF),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class CustomRadioButton extends StatelessWidget {
  const CustomRadioButton({
    required this.option,
    required this.selectedOption,
    required this.onSelected,
    this.suffixWidget,
    super.key,
  });

  final String option;
  final String? selectedOption;
  final ValueChanged<String> onSelected;
  final Widget? suffixWidget;

  @override
  Widget build(BuildContext context) {
    final isSelected = selectedOption == option;
    return GestureDetector(
      onTap: () => onSelected(option),
      child: Container(
        height: 38,
        width: double.maxFinite,
        margin: const EdgeInsets.only(bottom: 10),
        padding: const EdgeInsets.symmetric(horizontal: 10),
        child: Row(
          children: [
            Container(
              height: 24,
              width: 24,
              padding: const EdgeInsets.all(3),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                border: Border.all(
                  color:
                      isSelected ? context.colorScheme.primary : Colors.black,
                ),
              ),
              child: Container(
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: isSelected
                      ? context.colorScheme.primary
                      : Colors.transparent,
                ),
              ),
            ),
            const Space(20),
            Row(
              children: [
                Text(
                  option,
                  style: context.textTheme.bodyMedium,
                ),
                const Space(100),
                suffixWidget ?? Container(),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
