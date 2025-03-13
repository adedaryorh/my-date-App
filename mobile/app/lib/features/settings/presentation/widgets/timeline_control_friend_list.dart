import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class TimelineControlFriendsList extends StatelessWidget {
  const TimelineControlFriendsList({
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
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Row(
                children: [
                  Container(
                    height: 32,
                    width: 32,
                    decoration: const ShapeDecoration(
                      shape: CircleBorder(),
                      color: Color(0xffD9D9D9),
                    ),
                  ),
                  const Space(10),
                  Text(option),
                ],
              ),
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
            ],
          ),
          const Space(10),
          Container(
            height: 1,
            width: double.maxFinite,
            color: const Color(0xffE6E6E6),
          ),
          const Space(20),
        ],
      ),
    );
  }
}
