import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class LanguageOptionItem extends StatelessWidget {
  const LanguageOptionItem({
    required this.text,
    required this.flag,
    required this.tapped,
    required this.isSelected,
    super.key,
  });

  final String text;
  final String flag;
  final VoidCallback tapped;
  final bool isSelected;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: tapped,
      child: Container(
        width: double.maxFinite,
        margin: const EdgeInsets.only(bottom: 10),
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 5),
        height: 64,
        decoration: BoxDecoration(
          border: Border.all(
            color: isSelected
                ? context.colorScheme.primary
                : const Color(0xffEAEAEA),
          ),
          borderRadius: const BorderRadius.all(
            Radius.circular(16),
          ),
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(
              children: [
                Container(
                  alignment: Alignment.center,
                  height: 32,
                  width: 32,
                  decoration: const BoxDecoration(
                    shape: BoxShape.circle,
                    color: Color(0xffF3F6FB),
                  ),
                  child: Text(
                    flag,
                    style: context.textTheme.bodyLarge,
                  ),
                ),
                const SizedBox(
                  width: 10,
                ),
                Text(
                  text,
                  style: context.textTheme.bodyLarge,
                ),
              ],
            ),
            if (isSelected)
              Container(
                height: 20,
                width: 20,
                padding: const EdgeInsets.all(2),
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: context.colorScheme.primary,
                ),
                child: SvgPicture.asset(
                  AppAssets.checkIcon,
                  color: Colors.white,
                ),
              )
            else
              Container(),
          ],
        ),
      ),
    );
  }
}
